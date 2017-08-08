package cldr

import (
	"bytes"
	"fmt"
	"html/template"
	"reflect"
	"strings"
	"sync"
)

type pluralZero string
type pluralOne string
type pluralTwo string
type pluralFew string
type pluralMany string
type pluralOther string

// type pluralOther string

var parseTextCache = make(map[string]*template.Template)
var mutex sync.RWMutex
var parseFuncMap = template.FuncMap{
	"zero":  toPluralZero,
	"one":   toPluralOne,
	"two":   toPluralTwo,
	"few":   toPluralFew,
	"many":  toPluralMany,
	"other": toPluralOther,
}

func toPluralZero(text string) pluralZero   { return pluralZero(text) }
func toPluralOne(text string) pluralOne     { return pluralOne(text) }
func toPluralTwo(text string) pluralTwo     { return pluralTwo(text) }
func toPluralFew(text string) pluralFew     { return pluralFew(text) }
func toPluralMany(text string) pluralMany   { return pluralMany(text) }
func toPluralOther(text string) pluralOther { return pluralOther(text) }

// TODO: rename to T and add helper in Locale
func Parse(locale, text string, args ...interface{}) (r string, err error) {
	var data interface{}
	if len(args) > 0 {
		data = args[0]
	}

	mutex.RLock()
	tmpl := parseTextCache[text]
	mutex.RUnlock()

	p := parser{locale: locale, data: data}
	funcs := template.FuncMap{"p": p.parsePlural}
	for i := range args {
		arg := args[i]
		funcs[fmt.Sprintf("cldr_arg_%d", i+1)] = func() interface{} { return arg }
	}
	if tmpl == nil {
		for i := range args {
			text = strings.Replace(text, fmt.Sprintf("{{$%d}}", i+1), fmt.Sprintf("{{cldr_arg_%d}}", i+1), -1)
		}
		tmpl = template.New("").Funcs(parseFuncMap).Funcs(funcs)
		tmpl, err = tmpl.Parse(text)
		if err != nil {
			return
		}
		mutex.Lock()
		parseTextCache[text] = tmpl
		mutex.Unlock()
	} else {
		tmpl = tmpl.Funcs(funcs)
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, data); err != nil {
		return
	}
	r = buf.String()
	return
}

type parser struct {
	locale string
	data   interface{}
}

func (p parser) parsePlural(field string, rules ...interface{}) (r string, err error) {
	count, ok := getCount(p.data, field)
	if !ok {
		err = fmt.Errorf("can't find %s in %T", field, p.data)
		return
	}
	textMap := map[PluralRule]string{}
	for _, rule := range rules {
		switch x := rule.(type) {
		case pluralZero:
			textMap[PluralRuleZero] = string(x)
		case pluralOne:
			textMap[PluralRuleOne] = string(x)
		case pluralTwo:
			textMap[PluralRuleTwo] = string(x)
		case pluralFew:
			textMap[PluralRuleFew] = string(x)
		case pluralMany:
			textMap[PluralRuleMany] = string(x)
		case pluralOther:
			textMap[PluralRuleOther] = string(x)
		}
	}

	tmpl, err := template.New("").Parse(textMap[FindRule(p.locale, count)])
	if err != nil {
		return
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, p.data)
	r = buf.String()
	return
}

func getCount(v interface{}, field string) (interface{}, bool) {
	if len(field) == 0 {
		return nil, false
	}

	if field != "." {
		fieldParts := strings.Split(field, ".")
		fieldPartsLen := len(fieldParts)
		for i := 0; i < fieldPartsLen; i++ {
			fieldPart := fieldParts[i]
			rv := reflect.Indirect(reflect.ValueOf(v))
			switch rv.Kind() {
			case reflect.Struct:
				f := rv.FieldByName(fieldPart)
				if f.IsValid() {
					v = f.Interface()
					continue
				}
				return nil, false
			case reflect.Map:
				f := rv.MapIndex(reflect.ValueOf(fieldPart))
				if f.IsValid() {
					v = f.Interface()
					continue
				}
				return nil, false
			case reflect.Invalid:
				return nil, false
			}
		}
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Array, reflect.Chan, reflect.Slice:
		v = rv.Len()
	case reflect.Map, reflect.Struct:
		return nil, false
	}
	return v, true
}

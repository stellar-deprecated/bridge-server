package protocols

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

type Asset struct {
	Code   string `name:"asset_code"`
	Issuer string `name:"asset_issuer"`
}

type FormRequest struct{}

func (request *FormRequest) FromRequest(r *http.Request, destination interface{}) {
	rvalue := reflect.ValueOf(destination).Elem()
	typ := rvalue.Type()
	for i := 0; i < rvalue.NumField(); i++ {
		tag := typ.Field(i).Tag.Get("name")
		switch tag {
		case "":
			continue
		case "path":
			var path []Asset

			for i := 0; i < 5; i++ {
				codeFieldName := fmt.Sprintf("path[%d][asset_code]", i)
				issuerFieldName := fmt.Sprintf("path[%d][asset_issuer]", i)

				// If the element does not exist in PostForm break the loop
				if _, exists := r.PostForm[codeFieldName]; !exists {
					break
				}

				code := r.PostFormValue(codeFieldName)
				issuer := r.PostFormValue(issuerFieldName)

				if code == "" && issuer == "" {
					path = append(path, Asset{})
				} else {
					path = append(path, Asset{code, issuer})
				}
			}

			ptr := rvalue.Field(i).Addr().Interface().(*[]Asset)
			*ptr = path
		default:
			value := r.PostFormValue(tag)
			rvalue.Field(i).SetString(value)
		}
	}
	return
}

func (request *FormRequest) ToValues(object interface{}) (values url.Values) {
	values = make(map[string][]string)
	rvalue := reflect.ValueOf(object).Elem()
	typ := rvalue.Type()
	for i := 0; i < rvalue.NumField(); i++ {
		field := rvalue.Field(i)
		tag := typ.Field(i).Tag.Get("name")
		if tag == "" {
			continue
		}
		switch field.Interface().(type) {
		case string:
			value := rvalue.Field(i).String()
			if value == "" {
				continue
			}
			values.Set(tag, value)
		case []Asset:
			assets := rvalue.Field(i).Interface().([]Asset)
			for i, asset := range assets {
				values.Set(fmt.Sprintf("path[%d][asset_code]", i), asset.Code)
				values.Set(fmt.Sprintf("path[%d][asset_issuer]", i), asset.Issuer)
			}
		}
	}
	return
}

type SuccessResponse struct{}

func (response *SuccessResponse) HTTPStatus() int {
	return http.StatusOK
}

// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"go/format"

	"gopkg.in/yaml.v2"
)

// download github.com/vube/i18n first
func main() {
	rules := filepath.Join(strings.Split(os.Getenv("GOPATH"), ":")[0], "src/github.com/vube/i18n/data/rules")
	err := filepath.Walk(rules, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if info.IsDir() {
			return nil
		}
		in, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		var out yaml.MapSlice
		if err = yaml.Unmarshal(in, &out); err != nil {
			panic(err)
		}
		var plural string
		for _, item := range out {
			if item.Key == "plural" {
				plural = item.Value.(string)
			}
		}
		if plural == "" {
			fmt.Println("can't find plurals in", path)
			return nil
		}
		locale := strings.Replace(info.Name(), filepath.Ext(info.Name()), "", 1)
		locale = strings.Replace(locale, "-", "_", -1)
		dir := "resources/locales/" + locale
		if _, err := os.Stat(dir); err != nil {
			fmt.Println(dir, "doesn't exist")
			return nil
		}
		file, err := os.Create(dir + "/plural.go")
		if err != nil {
			panic(err)
		}
		rawcodes := fmt.Sprintf(`package %s

		var pluralRule = %q
		`, locale, plural)
		var codes []byte
		if codes, err = format.Source([]byte(rawcodes)); err != nil {
			panic(err)
		}
		if _, err := file.Write(codes); err != nil {
			panic(err)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

package db

import (
	"reflect"
	"sort"

	"github.com/jmoiron/sqlx/reflectx"
)

var mapper = reflectx.NewMapper("db")

type person struct {
	Name        string `db:"name"`
	HungerLevel string `db:"hunger_level"`

	SomethingIgnored int `db:"-"`
}

// columnsForStruct returns a slice of column names for the provided value
// (which should be a struct, a slice of structs).
func columnsForStruct(dest interface{}) []string {
	typ := reflect.TypeOf(dest)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() == reflect.Slice {
		typ = typ.Elem()
	}

	typmap := mapper.TypeMap(typ)

	var keys []string
	for k := range typmap.Names {
		keys = append(keys, k)
	}

	// Ensure keys are sorted.  keys is populated from a map, which has no
	// defined iteration order.  Different versions of go or different
	// architectures may cause non-deterministic results to occur (and in our CI
	// environment, they have).  To make testing easier, we sort the keys.
	sort.Strings(keys)

	return keys
}

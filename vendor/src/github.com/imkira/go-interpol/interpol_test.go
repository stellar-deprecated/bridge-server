package interpol

import (
	"errors"
	"io"
	"strings"
	"testing"
)

func toUpper(key string, w io.Writer) error {
	w.Write([]byte(strings.ToUpper(key)))
	return nil
}

func TestWithFunc(t *testing.T) {
	str, err := WithFunc("hello {test}!!!", toUpper)
	if err != nil {
		t.Fatal(err)
	}
	if str != "hello TEST!!!" {
		t.Errorf("Invalid string: %q", str)
	}
}

func TestWithFuncEscapeOpen(t *testing.T) {
	str, err := WithFunc("hello {{!!!{{", toUpper)
	if err != nil {
		t.Fatal(err)
	}
	if str != "hello {!!!{" {
		t.Errorf("Invalid string: %q", str)
	}
}

func TestWithFuncEscapeClose(t *testing.T) {
	str, err := WithFunc("hello }}!!!}}", toUpper)
	if err != nil {
		t.Fatal(err)
	}
	if str != "hello }!!!}" {
		t.Errorf("Invalid string: %q", str)
	}
}

func TestWithFuncEscapeOpenClose(t *testing.T) {
	str, err := WithFunc("hello }}!!!{{", toUpper)
	if err != nil {
		t.Fatal(err)
	}
	if str != "hello }!!!{" {
		t.Errorf("Invalid string: %q", str)
	}
}

func TestWithFuncUnexpectedClose(t *testing.T) {
	strs := []string{
		"}hello test!!!",
		"hello test}!!!",
		"hello test!!!}",
		"hello test!!!}{",
	}
	for _, s := range strs {
		str, err := WithFunc(s, toUpper)
		if len(str) != 0 || err != ErrUnexpectedClose {
			t.Fatal(err)
		}
	}
}

func TestWithFuncExpectingClose(t *testing.T) {
	strs := []string{
		"{hello test!!!",
		"hello {test!!!",
		"hello test!!!{",
	}
	for _, s := range strs {
		str, err := WithFunc(s, toUpper)
		if len(str) != 0 || err != ErrExpectingClose {
			t.Fatal(err)
		}
	}
}

func TestWithFuncFuncError(t *testing.T) {
	dummy := errors.New("dummy")
	f := func(key string, w io.Writer) error {
		return dummy
	}
	str, err := WithFunc("hello {test}!!!", f)
	if len(str) != 0 || err != dummy {
		t.Fatal(err)
	}
}

func TestWithMapKeyExists(t *testing.T) {
	m := map[string]string{
		"test": "World",
		"data": "!!!",
	}
	str, err := WithMap("hello {test}{data}", m)
	if err != nil {
		t.Fatal(err)
	}
	if str != "hello World!!!" {
		t.Errorf("Invalid string: %q", str)
	}
}

func TestWithMapKeyNotFound(t *testing.T) {
	m := map[string]string{
		"data": "World",
	}
	str, err := WithMap("hello {test}!!!", m)
	if len(str) != 0 || err != ErrKeyNotFound {
		t.Fatal(err)
	}
}

func TestWithMapNil(t *testing.T) {
	str, err := WithMap("hello {test}!!!", nil)
	if len(str) != 0 || err != ErrKeyNotFound {
		t.Fatal(err)
	}
}

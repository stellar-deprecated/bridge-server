package scs_test

import (
	"testing"

	"github.com/alexedwards/scs/engine/memstore"
	"github.com/qor/session/scs"
	"github.com/qor/session/test"
)

func TestAll(t *testing.T) {
	engine := memstore.New(0)
	manager := scs.New(engine)
	test.TestAll(manager, t)
}

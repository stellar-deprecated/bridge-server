package filesystem

import (
	"testing"

	"github.com/qor/oss/tests"
)

func TestAll(t *testing.T) {
	fileSystem := New("/tmp")
	tests.TestAll(fileSystem, t)
}

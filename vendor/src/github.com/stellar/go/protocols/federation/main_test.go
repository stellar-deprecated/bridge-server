package federation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	var m Memo

	m = Memo{"123"}
	value, err := json.Marshal(m)
	assert.NoError(t, err)
	assert.Equal(t, "123", string(value))

	m = Memo{"Test"}
	value, err = json.Marshal(m)
	assert.NoError(t, err)
	assert.Equal(t, `"Test"`, string(value))
}

func TestUnmarshal(t *testing.T) {
	var m Memo

	err := json.Unmarshal([]byte("123"), &m)
	assert.NoError(t, err)
	assert.Equal(t, "123", m.Value)

	err = json.Unmarshal([]byte(`"123"`), &m)
	assert.NoError(t, err)
	assert.Equal(t, "123", m.Value)

	err = json.Unmarshal([]byte(`"Test"`), &m)
	assert.NoError(t, err)
	assert.Equal(t, "Test", m.Value)

	err = json.Unmarshal([]byte("-123"), &m)
	assert.Error(t, err)
}

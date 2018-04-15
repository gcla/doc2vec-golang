package common

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	data := []string{"a", "b", "c"}
	a := NewArrayStringProvider(data)
	var _ IStringProvider = a

	assert.Equal(t, true, a.More(), "foo")
	assert.Equal(t, "a", a.Next(), "foo")
	assert.Equal(t, true, a.More(), "foo")
	assert.Equal(t, "b", a.Next(), "foo")
	assert.Equal(t, true, a.More(), "foo")
	assert.Equal(t, "c", a.Next(), "foo")
	assert.Equal(t, false, a.More(), "foo")
}

func Test2(t *testing.T) {
	data := strings.Join([]string{
		"ep1\taa bb cc",
		"ep2\tdd ee ff",
	}, "\n")

	a := NewModelFileDataProvider(strings.NewReader(data))
	var _ IModelDataProvider = a
	var words IStringProvider

	assert.Equal(t, true, a.More(), "foo")
	ep1 := a.Next()
	assert.Equal(t, "ep1", ep1.Name(), "foo")
	words = ep1.Words()
	assert.Equal(t, true, words.More(), "foo")
	assert.Equal(t, "aa", words.Next(), "foo")
	assert.Equal(t, true, words.More(), "foo")
	assert.Equal(t, "bb", words.Next(), "foo")
	assert.Equal(t, true, words.More(), "foo")
	assert.Equal(t, "cc", words.Next(), "foo")
	assert.Equal(t, false, words.More(), "foo")

	assert.Equal(t, true, a.More(), "foo")
	ep2 := a.Next()
	assert.Equal(t, "ep2", ep2.Name(), "foo")
	words = ep2.Words()
	assert.Equal(t, true, words.More(), "foo")
	assert.Equal(t, "dd", words.Next(), "foo")
	assert.Equal(t, true, words.More(), "foo")
	assert.Equal(t, "ee", words.Next(), "foo")
	assert.Equal(t, true, words.More(), "foo")
	assert.Equal(t, "ff", words.Next(), "foo")
	assert.Equal(t, false, words.More(), "foo")

	assert.Equal(t, false, a.More(), "foo")

	a.Reset()
	assert.Equal(t, true, a.More(), "foo")
	ep3 := a.Next()
	assert.Equal(t, "ep1", ep3.Name(), "foo")
	words = ep3.Words()
	assert.Equal(t, true, words.More(), "foo")
	assert.Equal(t, "aa", words.Next(), "foo")
	assert.Equal(t, true, words.More(), "foo")
	assert.Equal(t, "bb", words.Next(), "foo")
	assert.Equal(t, true, words.More(), "foo")
	assert.Equal(t, "cc", words.Next(), "foo")
	assert.Equal(t, false, words.More(), "foo")
}

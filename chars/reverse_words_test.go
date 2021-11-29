package chars

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestReverseWords(t *testing.T) {
	s := "the sky is blue"
	assert.Equal(t, "blue is sky the", ReverseWords(s))
}

func TestReverseWords2(t *testing.T) {
	s := "  hello world  "
	assert.Equal(t, "world hello", ReverseWords(s))
}

func TestReverseWords3(t *testing.T) {
	s := "a good   example"
	assert.Equal(t, "example good a", ReverseWords(s))
}

func TestReverseWords4(t *testing.T) {
	s := "  Bob    Loves  Alice   "
	assert.Equal(t, "Alice Loves Bob", ReverseWords(s))
}

func TestReverseWords5(t *testing.T) {
	s := "Alice does not even like bob"
	assert.Equal(t, "bob like even not does Alice", ReverseWords(s))
}

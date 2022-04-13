package hot_one_hundred

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	s := "()"
	assert.Equal(t, true, IsValid(s))
}

func TestIsValid2(t *testing.T) {
	s := "()[]{}"
	assert.Equal(t, true, IsValid(s))
}

func TestIsValid3(t *testing.T) {
	s := "([)]"
	assert.Equal(t, false, IsValid(s))
}

func TestIsValid4(t *testing.T) {
	s := "{[]}"
	assert.Equal(t, true, IsValid(s))
}

//"(){}}{"

func TestIsValid5(t *testing.T) {
	s := "(){}}{"
	assert.Equal(t, false, IsValid(s))
}

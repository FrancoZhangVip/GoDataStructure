package chars

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestReplaceSpace(t *testing.T) {
	s := "We are happy."
	res := ReplaceSpace(s)
	assert.Equal(t, "We%20are%20happy.", res)
}

func TestReplaceSpace2(t *testing.T) {
	s := "   "
	res := ReplaceSpace(s)
	assert.Equal(t, "%20%20%20", res)
}

func TestReplaceSpace3(t *testing.T) {
	s := " "
	res := ReplaceSpace(s)
	assert.Equal(t, "%20", res)
}

func TestReplaceSpace4(t *testing.T) {
	s := ""
	res := ReplaceSpace(s)
	assert.Equal(t, "", res)
}

package sword_towards_offer

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestReplaceSpace(t *testing.T) {
	s := "We are happy."
	assert.Equal(t, "We%20are%20happy.", ReplaceSpace(s))
}

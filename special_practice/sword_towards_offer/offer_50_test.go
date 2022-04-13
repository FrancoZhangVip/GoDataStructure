package sword_towards_offer

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	s := "abaccdeff"
	assert.Equal(t, uint8('b'), FirstUniqChar(s))
}

func TestFirstUniqChar2(t *testing.T) {
	s := ""
	assert.Equal(t, uint8(' '), FirstUniqChar(s))
}

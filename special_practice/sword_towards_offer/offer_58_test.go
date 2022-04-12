package sword_towards_offer

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestReverseLeftWords(t *testing.T) {
	s := "abcdefg"
	k := 2
	assert.Equal(t, "cdefgab", ReverseLeftWords(s, k))
}

func TestReverseLeftWords2(t *testing.T) {
	s := "lrloseumgh"
	k := 6
	assert.Equal(t, "umghlrlose", ReverseLeftWords(s, k))
}

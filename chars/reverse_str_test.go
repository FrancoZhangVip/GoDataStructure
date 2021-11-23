package chars

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestReverseStr(t *testing.T) {
	s := "abcdefg"
	k := 2
	assert.Equal(t, "bacdfeg", ReverseStr(s, k))
}

func TestReverseStr2(t *testing.T) {
	s := "abcdefgh"
	k := 3
	assert.Equal(t, "cbadefhg", ReverseStr(s, k))
}

func TestReverseStr3(t *testing.T) {
	s := "abcd"
	k := 2
	assert.Equal(t, "bacd", ReverseStr(s, k))
}

func TestReverseStr4(t *testing.T) {
	s := "abcdefg"
	k := 8
	assert.Equal(t, "gfedcba", ReverseStr(s, k))
}

//"hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl"
//39
func TestReverseStr5(t *testing.T) {
	s := "hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl"
	k := 39
	assert.Equal(t, "fdcqkmxwholhytmhafpesaentdvxginrjlyqzyhehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqllgsqddebemjanqcqnfkjmi", ReverseStr(s, k))

	//fdcqkmxwholhytmhafpesaentdvxginrjlyqzyhehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqllgsqddebemjanqcqnfkjmi
	//fdcqkmxwholhytmhafpesaentdvxginrjlyqzyhehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqllgskfnqcqnajmebeddqjmi
}

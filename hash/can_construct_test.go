package hash

import "testing"

func TestCanConstruct(t *testing.T) {
	ransomNote := "a"
	magazine := "b"
	println(CanConstruct(ransomNote, magazine))
}

func TestCanConstruct2(t *testing.T) {
	ransomNote := "aa"
	magazine := "ab"
	println(CanConstruct(ransomNote, magazine))
}

func TestCanConstruct3(t *testing.T) {
	ransomNote := "aa"
	magazine := "aab"
	println(CanConstruct(ransomNote, magazine))
}

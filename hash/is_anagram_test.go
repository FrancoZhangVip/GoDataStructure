package hash

import "testing"

func TestIsAnagram(t *testing.T) {
	s := "anagram"
	te := "nagaram"

	println(IsAnagram(s, te))
}

func TestIsAnagram2(t *testing.T) {
	s := "rat"
	te := "car"

	println(IsAnagram(s, te))
}

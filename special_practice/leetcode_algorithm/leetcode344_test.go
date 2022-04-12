package leetcode_algorithm

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	ReverseString(s)
	fmt.Println(string(s))
}

func TestReverseString2(t *testing.T) {
	s := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	ReverseString(s)
	fmt.Println(string(s))
}

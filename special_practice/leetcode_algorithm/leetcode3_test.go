package leetcode_algorithm

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	assert.Equal(t, 3, LengthOfLongestSubstring(s))
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	s := "bbbbb"
	assert.Equal(t, 1, LengthOfLongestSubstring(s))
}

func TestLengthOfLongestSubstring3(t *testing.T) {
	s := "pwwkew"
	assert.Equal(t, 3, LengthOfLongestSubstring(s))
}

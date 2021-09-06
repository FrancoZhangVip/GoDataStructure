package algorithm_init

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := ""
	println("Test 1 is: ", LengthOfLongestSubstring(s)) //0
	assert.Equal(t, 0, LengthOfLongestSubstring(s))

	s = "abcabcbb"
	println("Test 2 is: ", LengthOfLongestSubstring(s)) //3
	assert.Equal(t, 3, LengthOfLongestSubstring(s))

	s = "bbbbb"
	println("Test 3 is: ", LengthOfLongestSubstring(s)) //1
	assert.Equal(t, 1, LengthOfLongestSubstring(s))

	s = "pwwkew"
	println("Test 4 is: ", LengthOfLongestSubstring(s)) //3
	assert.Equal(t, 3, LengthOfLongestSubstring(s))

	s = " "
	println("Test 5 is: ", LengthOfLongestSubstring(s)) //1
	assert.Equal(t, 1, LengthOfLongestSubstring(s))

	s = "aa"
	println("Test 6 is: ", LengthOfLongestSubstring(s)) //1
	assert.Equal(t, 1, LengthOfLongestSubstring(s))

	s = "aab"
	println("Test 7 is: ", LengthOfLongestSubstring(s)) //2
	assert.Equal(t, 2, LengthOfLongestSubstring(s))
}

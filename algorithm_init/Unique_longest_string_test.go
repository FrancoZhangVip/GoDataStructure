package algorithm_init

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	s := ""
	println("Test 1 is: ", LengthOfLongestSubstring(s)) //0

	s = "abcabcbb"
	println("Test 2 is: ", LengthOfLongestSubstring(s)) //3

	s = "bbbbb"
	println("Test 3 is: ", LengthOfLongestSubstring(s)) //1

	s = "pwwkew"
	println("Test 4 is: ", LengthOfLongestSubstring(s)) //3

	s = " "
	println("Test 5 is: ", LengthOfLongestSubstring(s)) //1

	s = "aa"
	println("Test 6 is: ", LengthOfLongestSubstring(s)) //1

	s = "aab"
	println("Test 7 is: ", LengthOfLongestSubstring(s)) //2
}

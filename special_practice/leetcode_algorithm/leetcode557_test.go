package leetcode_algorithm

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestReverseWords(t *testing.T) {
	s := "Let's take LeetCode contest"
	assert.Equal(t, "s'teL ekat edoCteeL tsetnoc", ReverseWords(s))
}

func TestReverseWords2(t *testing.T) {
	s := "God Ding"
	assert.Equal(t, "doG gniD", ReverseWords(s))
}

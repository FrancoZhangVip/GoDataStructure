package leetcode_algorithm

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	assert.Equal(t, 2, SearchInsert(nums, target))
}

func TestSearchInsert2(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 2
	assert.Equal(t, 1, SearchInsert(nums, target))
}

func TestSearchInsert3(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 7
	assert.Equal(t, 4, SearchInsert(nums, target))
}

package leetcode_algorithm

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	assert.Equal(t, 4, Search(nums, target))
}

func TestSearch2(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	assert.Equal(t, -1, Search(nums, target))
}

/*
[-1,0,3,5,9,12]
13
*/
func TestSearch3(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 13
	assert.Equal(t, -1, Search(nums, target))
}

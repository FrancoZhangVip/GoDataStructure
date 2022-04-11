package leetcode_algorithm

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	nums := []int{-4, -1, 0, 3, 10}
	res := SortedSquares(nums)
	assert.Equal(t, 0, res[0])
	assert.Equal(t, 1, res[1])
	assert.Equal(t, 9, res[2])
	assert.Equal(t, 16, res[3])
	assert.Equal(t, 100, res[4])
}

func TestSortedSquares2(t *testing.T) {
	nums := []int{-7, -3, 2, 3, 11}
	res := SortedSquares(nums)
	assert.Equal(t, 4, res[0])
	assert.Equal(t, 9, res[1])
	assert.Equal(t, 9, res[2])
	assert.Equal(t, 49, res[3])
	assert.Equal(t, 121, res[4])
}

func TestSortedSquares3(t *testing.T) {
	nums := []int{-7, -3}
	res := SortedSquares(nums)
	assert.Equal(t, 9, res[0])
	assert.Equal(t, 49, res[1])
}

func TestSortedSquares4(t *testing.T) {
	nums := []int{1}
	res := SortedSquares(nums)
	assert.Equal(t, 1, res[0])
}

func TestSortedSquares5(t *testing.T) {
	nums := []int{-4, -4, -3}
	res := SortedSquares(nums)
	assert.Equal(t, 9, res[0])
	assert.Equal(t, 16, res[1])
	assert.Equal(t, 16, res[2])
}

package array

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestBinary_Search(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	assert.Equal(t, 4, Binary_Search(nums, 9))
	assert.Equal(t, -1, Binary_Search(nums, 2))
	assert.Equal(t, 0, Binary_Search(nums, -1))
	assert.Equal(t, 5, Binary_Search(nums, 12))
}

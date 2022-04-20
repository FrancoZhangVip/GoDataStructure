package hot_one_hundred

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	assert.Equal(t, 6, MaxSubArray(nums))
}

func TestMaxSubArray2(t *testing.T) {
	nums := []int{1}
	assert.Equal(t, 1, MaxSubArray(nums))
}

func TestMaxSubArray3(t *testing.T) {
	nums := []int{5, 4, -1, 7, 8}
	assert.Equal(t, 23, MaxSubArray(nums))
}

package leetcode_algorithm

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	numbers := []int{2, 7, 11, 15}
	target := 9

	res := TwoSum(numbers, target)
	assert.Equal(t, 2, len(res))
	assert.Equal(t, 1, res[0])
	assert.Equal(t, 2, res[1])
}

func TestTwoSum2(t *testing.T) {
	numbers := []int{2, 3, 4}
	target := 6

	res := TwoSum(numbers, target)
	assert.Equal(t, 2, len(res))
	assert.Equal(t, 1, res[0])
	assert.Equal(t, 3, res[1])
}

func TestTwoSum3(t *testing.T) {
	numbers := []int{-1, 0}
	target := -1

	res := TwoSum(numbers, target)
	assert.Equal(t, 2, len(res))
	assert.Equal(t, 1, res[0])
	assert.Equal(t, 2, res[1])
}

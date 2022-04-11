package leetcode_algorithm

import "testing"

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	MoveZeroes(nums)
	ShowNums(nums)
}

func TestMoveZeroes2(t *testing.T) {
	nums := []int{0}
	MoveZeroes(nums)
	ShowNums(nums)
}

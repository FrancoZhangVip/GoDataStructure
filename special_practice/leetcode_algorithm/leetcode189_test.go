package leetcode_algorithm

import (
	"fmt"
	"testing"
)

func ShowNums(nums []int) {
	for _, data := range nums {
		fmt.Print(data, "\t")
	}
	fmt.Println()
}

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	Rotate(nums, k)
	ShowNums(nums)
}

func TestRotate2(t *testing.T) {
	nums := []int{-1, -100, 3, 99}
	k := 2
	Rotate(nums, k)
	ShowNums(nums)
}

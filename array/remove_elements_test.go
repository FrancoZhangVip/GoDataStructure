package array

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	show(nums, val)

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	show(nums, val)

	nums = []int{}
	val = 2
	show(nums, val)

	nums = []int{2, 2, 2, 2, 2}
	val = 2
	show(nums, val)

	nums = []int{2}
	val = 2
	show(nums, val)

	nums = []int{3}
	val = 2
	show(nums, val)
}

func show(nums []int, val int) {
	templen := RemoveElement(nums, val)
	for i := 0; i < templen; i++ {
		fmt.Print(nums[i], "\t")
	}
	fmt.Println()
}

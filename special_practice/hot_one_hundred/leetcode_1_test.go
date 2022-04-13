package hot_one_hundred

import (
	"fmt"
	"testing"
)

func ShowArray(nums []int) {
	for _, data := range nums {
		fmt.Print(data, "\t")
	}
	fmt.Println()
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	res := TwoSum(nums, target)
	ShowArray(res)
}

func TestTwoSum2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6

	res := TwoSum(nums, target)
	ShowArray(res)
}

func TestTwoSum3(t *testing.T) {
	nums := []int{3, 3}
	target := 6

	res := TwoSum(nums, target)
	ShowArray(res)
}

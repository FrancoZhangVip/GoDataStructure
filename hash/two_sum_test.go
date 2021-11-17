package hash

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	res := TwoSum(nums, target)

	ShowArrayList(res)
}

func TestTwoSum2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6

	res := TwoSum(nums, target)

	ShowArrayList(res)
}

func TestTwoSum3(t *testing.T) {
	nums := []int{3, 3}
	target := 6

	res := TwoSum(nums, target)

	ShowArrayList(res)
}

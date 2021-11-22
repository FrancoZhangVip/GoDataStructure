package hash

import "testing"

func TestFourSum(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	res := FourSum(nums, 0)
	ShowArrayNow(res)
}

func TestFourSum2(t *testing.T) {
	nums := []int{2, 2, 2, 2, 2}
	res := FourSum(nums, 8)
	ShowArrayNow(res)
}

package hash

import "testing"

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := ThreeSum(nums)
	ShowArrayNow(res)
}

func TestThreeSum2(t *testing.T) {
	nums := []int{0, 0, 0, 0}
	res := ThreeSum(nums)
	ShowArrayNow(res)
}

func ShowArrayNow(nums [][]int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			print(nums[i][j], "\t")
		}
		println()
	}

}

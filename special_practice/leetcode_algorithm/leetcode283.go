package leetcode_algorithm

func MoveZeroes(nums []int) {
	p, q := 0, 0
	for p < len(nums) && q < len(nums) {
		if nums[q] != 0 {
			nums[p] = nums[q]
			p++
			q++
		} else {
			q++
		}
	}
	for i := p; i < len(nums); i++ {
		nums[i] = 0
	}
}

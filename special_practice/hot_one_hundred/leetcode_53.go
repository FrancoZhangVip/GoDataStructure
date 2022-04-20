package hot_one_hundred

func MaxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

package hot_one_hundred

func FindDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		index := (nums[i] % len(nums)) - 1
		if index == -1 {
			index = len(nums) - 1
		}
		nums[index] += len(nums)
	}
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] <= len(nums) && nums[i] >= 1 {
			res = append(res, i+1)
		}
	}
	return res
}

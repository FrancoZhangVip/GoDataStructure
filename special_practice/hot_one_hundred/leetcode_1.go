package hot_one_hundred

func TwoSum(nums []int, target int) []int {
	checker := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := checker[target-nums[i]]; !ok {
			checker[nums[i]] = i
		} else {
			res := make([]int, 0)
			res = append(res, checker[target-nums[i]], i)
			return res
		}
	}
	return make([]int, 0)
}

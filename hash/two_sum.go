package hash

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

func TwoSum(nums []int, target int) []int {
	map_checker := make(map[int]int)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		_, ok := map_checker[target-nums[i]]
		if ok {
			res = append(res, i, map_checker[target-nums[i]])
			return res
		} else {
			map_checker[nums[i]] = i
		}
	}
	return res
}

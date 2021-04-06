package algorithm_init

/*
subsets
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
*/

func Subsets(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) == 0 {
		return result
	}

	//init
	result = append(result, make([]int, 0))
	result = append(result, append(make([]int, 0), nums[0]))

	for i := 1; i < len(nums); i++ {
		tmp := make([][]int, 0)
		for j := 0; j < len(result); j++ {
			//copy
			target := make([]int, 0)
			for k := 0; k < len(result[j]); k++ {
				target = append(target, result[j][k])
			}

			//add current num
			target = append(target, nums[i])
			tmp = append(tmp, target)
		}

		result = append(result, tmp...)
	}

	return result
}

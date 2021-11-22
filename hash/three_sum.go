package hash

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
注意： 答案中不可以包含重复的三元组。

示例：
给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为： [ [-1, 0, 1], [-1, -1, 2] ]
*/

func ThreeSum(nums []int) [][]int {
	res := make([][]int, 0)

	if len(nums) < 3 {
		return res
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] < nums[j] {
				temp := nums[j]
				nums[j] = nums[i]
				nums[i] = temp
			}
		}
	}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		slow := i + 1
		quick := len(nums) - 1
		for true {
			if !(i < slow && slow < quick) {
				break
			}
			calcu := nums[i] + nums[slow] + nums[quick]
			num_slow := nums[slow]
			num_quick := nums[quick]
			if calcu == 0 {
				resItem := make([]int, 0)
				resItem = append(resItem, nums[i], nums[slow], nums[quick])
				res = append(res, resItem)
				for slow < quick && nums[slow] == num_slow {
					slow++
				}
				for slow < quick && nums[quick] == num_quick {
					quick--
				}

			} else if calcu < 0 {
				slow++
			} else {
				quick--
			}
		}
	}
	return res
}

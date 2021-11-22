package hash

/*
题意：给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，
使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：
答案中不可以包含重复的四元组。

示例：
给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
满足要求的四元组集合为： [ [-1, 0, 0, 1], [-2, -1, 1, 2], [-2, 0, 0, 2] ]
*/

func FourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 4 {
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

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l := j + 1
			r := len(nums) - 1
			for true {
				if !(j < l && l < r) {
					break
				}
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				num_l := nums[l]
				num_r := nums[r]
				if sum == target {
					res_item := make([]int, 0)
					res_item = append(res_item, nums[i], nums[j], nums[l], nums[r])
					res = append(res, res_item)
					for l < r && nums[l] == num_l {
						l++
					}
					for l < r && nums[r] == num_r {
						r--
					}
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}

	}
	return res
}

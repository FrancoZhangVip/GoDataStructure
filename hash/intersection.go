package hash

/*
给定两个数组，编写一个函数来计算它们的交集。

示例 1：
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
示例 2：

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]


链接：https://leetcode-cn.com/problems/intersection-of-two-arrays

*/

func Intersection(nums1 []int, nums2 []int) []int {
	checker_map := make(map[int]int)
	res := make([]int, 0)

	for _, num := range nums1 {
		if checker_map[num] == 0 {
			checker_map[num] = 1
		}
	}

	for _, num := range nums2 {
		if checker_map[num] == 0 {
			checker_map[num] = 2
		} else if checker_map[num] == 1 {
			checker_map[num] = 3
		}
	}

	for k, v := range checker_map {
		if v == 3 {
			res = append(res, k)
		}
	}

	return res

}

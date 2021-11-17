package hash

/*
给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。
为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -2^28 到 2^28 - 1 之间，最终结果不会超过 2^31 - 1 。

例如:
输入: A = [ 1, 2] B = [-2,-1] C = [-1, 2] D = [ 0, 2] 输出: 2 解释: 两个元组如下:
(0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
(1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
*/

func FourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	counter := 0
	map_checker := make(map[int]int)
	for _, a := range nums1 {
		for _, b := range nums2 {
			map_checker[a+b]++
		}
	}

	for _, c := range nums3 {
		for _, d := range nums4 {
			if map_checker[-(c+d)] > 0 {
				counter = counter + map_checker[-(c+d)]
			}
		}
	}
	return counter
}

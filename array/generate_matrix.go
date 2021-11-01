package array

/*
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

示例 1：
|1|2|3|
|8|9|4|
|7|6|5|
输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]
示例 2：

输入：n = 1
输出：[[1]]


提示：

1 <= n <= 20

https://leetcode-cn.com/problems/spiral-matrix-ii

*/

func GenerateMatrix(n int) [][]int {

	//init matrix
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	raw := 0
	col := 0
	counter := 1
	len := n
	for counter <= n*n {

		for i := 0; i < len; i++ {
			res[raw][col] = counter
			col++
			counter++
		}
		col--
		raw++
		for i := 0; i < len-1; i++ {
			res[raw][col] = counter
			raw++
			counter++
		}
		raw--
		col--
		for i := 0; i < len-1; i++ {
			res[raw][col] = counter
			col--
			counter++
		}
		col++
		raw--
		for i := 0; i < len-2; i++ {
			res[raw][col] = counter
			raw--
			counter++
		}
		raw++
		col++

		len = len - 2
	}
	return res
}

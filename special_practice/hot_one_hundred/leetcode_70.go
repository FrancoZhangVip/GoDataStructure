package hot_one_hundred

func ClimbStairs(n int) int {
	dict := make([]int, 0)
	dict = append(dict, 1)
	dict = append(dict, 2)
	for i := 3; i <= n; i++ {
		val := dict[i-3] + dict[i-2]
		dict = append(dict, val)
	}
	return dict[n-1]
}

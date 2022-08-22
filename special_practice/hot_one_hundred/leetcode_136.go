package hot_one_hundred

func SingleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single = single ^ num
	}
	return single
}

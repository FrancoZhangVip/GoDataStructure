package sword_towards_offer

func FindRepeatNumber(nums []int) int {
	checker := make(map[int]int)

	for _, data := range nums {
		if checker[data] == 1 {
			return data
		} else {
			checker[data] = 1
		}
	}
	return -1
}

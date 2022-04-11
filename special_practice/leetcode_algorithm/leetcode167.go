package leetcode_algorithm

func TwoSum(numbers []int, target int) []int {
	p, q := 0, len(numbers)-1
	for true {
		if q <= p {
			break
		}
		if numbers[p]+numbers[q] == target {
			return []int{p + 1, q + 1}
		} else if numbers[p]+numbers[q] < target {
			p++
		} else if numbers[p]+numbers[q] > target {
			q--
		}
	}
	return []int{-1, -1}
}

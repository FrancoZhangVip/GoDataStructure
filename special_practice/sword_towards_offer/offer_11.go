package sword_towards_offer

import "math"

func MinArray(numbers []int) int {
	p, q := 0, len(numbers)-1

	for p < q-1 {
		mid := (p + q) / 2
		if numbers[mid] > numbers[q] {
			p = mid
		} else if numbers[mid] < numbers[q] {
			q = mid
		} else {
			q--
		}
	}

	return int(math.Min(float64(numbers[p]), float64(numbers[q])))
}

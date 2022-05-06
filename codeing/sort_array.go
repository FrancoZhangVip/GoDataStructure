package codeing

import (
	"sort"
)

func SortArray(data []int) []int {
	res1 := make([]int, 0)
	for _, item := range data {
		if item > 0 {
			res1 = append(res1, item)
		}
	}
	res2 := make([]int, 0)
	for _, item := range data {
		if item < 0 {
			res1 = append(res2, item)
		}
	}
	sort.Ints(res1)
	sort.Ints(res2)
	for i := len(res2) - 1; i > -1; i-- {
		res1 = append(res1, res2[i])
	}

	return res2
}

func SortArray2(data []int) []int {
	p, q := 0, len(data)-1
	for p >= q {
		if data[p] > 0 {
			p++
			continue
		}
		if data[q] < 0 {
			q--
			continue
		}
		data[p], data[q] = data[q], data[p]
	}
	return data
}

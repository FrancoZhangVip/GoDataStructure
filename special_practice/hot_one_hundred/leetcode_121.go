package hot_one_hundred

import "math"

func MaxProfit(prices []int) int {
	minPrice := math.MaxInt64
	maxProfit := 0
	for _, val := range prices {
		if val < minPrice {
			minPrice = val
		}
		if val-minPrice > maxProfit {
			maxProfit = val - minPrice
		}
	}
	return maxProfit
}

package array

import (
	"fmt"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	nums := []int{-4, -1, 0, 3, 10}
	res := SortedSquares(nums)
	SortedSquaresShow(res)

	nums = []int{-4}
	res = SortedSquares(nums)
	SortedSquaresShow(res)

	nums = []int{}
	res = SortedSquares(nums)
	SortedSquaresShow(res)

	nums = []int{-7, -3, 2, 3, 11}
	res = SortedSquares(nums)
	SortedSquaresShow(res)
}

func SortedSquaresShow(nums []int) {
	for _, val := range nums {
		fmt.Print(val, "\t")

	}
	fmt.Println()
}

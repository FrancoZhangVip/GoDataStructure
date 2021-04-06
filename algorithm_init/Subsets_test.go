package algorithm_init

import (
	"fmt"
	"testing"
)

func Show(result [][]int) {
	for _, arraySingle := range result {
		fmt.Println(arraySingle, "\t")
	}
}

func TestSubsets(t *testing.T) {
	nums := make([]int, 0)
	nums = append(nums, 1, 2, 3)

	result := Subsets(nums)

	Show(result)
}

func TestSubsets_Lot(t *testing.T) {
	nums := make([]int, 0)
	nums = append(nums, 1, 2, 3, 4, 5, 0, 9, 8, 7, 6)

	result := Subsets(nums)

	Show(result)
}

package array

import "testing"

func TestGenerateMatrix(t *testing.T) {
	res := GenerateMatrix(5)
	ShowArray(res)
}

func ShowArray(res [][]int) {
	for _, val := range res {
		for _, value := range val {
			print(value, "\t")
		}
		println()
	}
}

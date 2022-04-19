package leetcode_algorithm

import (
	"fmt"
	"testing"
)

func ShowArray(image [][]int) {
	for _, row := range image {
		for _, col := range row {
			fmt.Print(col, "\t")
		}
		fmt.Println()
	}
	fmt.Println()
}

func TestFloodFill(t *testing.T) {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr := 1
	sc := 1
	newColor := 2
	res := FloodFill(image, sr, sc, newColor)
	ShowArray(res)
}

func TestFloodFill2(t *testing.T) {
	image := [][]int{{0, 0, 0}, {0, 0, 0}}
	sr := 0
	sc := 0
	newColor := 2
	res := FloodFill(image, sr, sc, newColor)
	ShowArray(res)
}

package leetcode_algorithm

func MaxAreaOfIsland(grid [][]int) int {
	max := 0
	for i, row := range grid {
		for j, col := range row {
			if col == 1 {
				size := floodFill(grid, i, j, 2, 0)
				if max < size {
					max = size
				}
			}
		}
	}
	return max
}

func floodFill(image [][]int, sr int, sc int, newColor int, counter int) int {
	oldColor := image[sr][sc]
	image[sr][sc] = newColor
	counter++
	if sr-1 >= 0 && oldColor == image[sr-1][sc] {
		counter = floodFill(image, sr-1, sc, newColor, counter)
	}
	if sc-1 >= 0 && oldColor == image[sr][sc-1] {
		counter = floodFill(image, sr, sc-1, newColor, counter)
	}
	if sr+1 < len(image) && oldColor == image[sr+1][sc] {
		counter = floodFill(image, sr+1, sc, newColor, counter)
	}
	if sc+1 < len(image[0]) && oldColor == image[sr][sc+1] {
		counter = floodFill(image, sr, sc+1, newColor, counter)
	}
	return counter
}

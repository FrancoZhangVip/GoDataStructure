package leetcode_algorithm

func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	oldColor := image[sr][sc]
	image[sr][sc] = newColor
	if sr-1 >= 0 && oldColor == image[sr-1][sc] {
		FloodFill(image, sr-1, sc, newColor)
	}
	if sc-1 >= 0 && oldColor == image[sr][sc-1] {
		FloodFill(image, sr, sc-1, newColor)
	}
	if sr+1 < len(image) && oldColor == image[sr+1][sc] {
		FloodFill(image, sr+1, sc, newColor)
	}
	if sc+1 < len(image[0]) && oldColor == image[sr][sc+1] {
		FloodFill(image, sr, sc+1, newColor)
	}
	return image
}

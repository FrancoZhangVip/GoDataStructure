package sword_towards_offer

func FindNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, col := 0, len(matrix[0])-1
	for row != len(matrix) && col != -1 {
		item := matrix[row][col]
		if item == target {
			return true
		}
		if item > target {
			col--
		}
		if item < target {
			row++
		}
	}
	return false
}

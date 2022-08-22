package hot_one_hundred

func HammingDistance(x int, y int) int {
	return CountDistance(x ^ y)
}

func CountDistance(num int) (res int) {
	shang, yu := num, 0
	for num != 0 {
		yu = num % 2
		shang = num / 2
		if yu == 1 {
			res++
		}
		num = shang
	}
	return res
}

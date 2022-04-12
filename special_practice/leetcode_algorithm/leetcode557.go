package leetcode_algorithm

func ReverseWords(s string) string {
	data := []byte(s)
	p, q := 0, 0
	for true {
		if p >= len(data) || q >= len(data) {
			break
		}
		for true {
			if q < len(data) && data[q] != ' ' {
				q++
			} else {
				break
			}
		}

		for i := p; i < (p+q)/2; i++ {
			data[i], data[q-1-i+p] = data[q-1-i+p], data[i]
		}

		for true {
			if q < len(data) && data[q] == ' ' {
				q++
			} else {
				break
			}
		}

		p = q
	}
	return string(data)
}

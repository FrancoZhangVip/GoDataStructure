package hot_one_hundred

func IsValid(s string) bool {
	data := []byte(s)
	checker := make([]byte, len(s))
	checkerSize := 0

	for i := 0; i < len(data); i++ {
		if data[i] == '(' || data[i] == '{' || data[i] == '[' {
			checker[checkerSize] = data[i]
			checkerSize++
		} else {
			switch data[i] {
			case ')':
				if checkerSize-1 < 0 {
					return false
				}
				if checker[checkerSize-1] == '(' {
					checkerSize--
				} else {
					return false
				}
			case ']':
				if checkerSize-1 < 0 {
					return false
				}
				if checker[checkerSize-1] == '[' {
					checkerSize--
				} else {
					return false
				}
			case '}':
				if checkerSize-1 < 0 {
					return false
				}
				if checker[checkerSize-1] == '{' {
					checkerSize--
				} else {
					return false
				}
			}
		}
	}
	if checkerSize == 0 {
		return true
	}
	return false
}

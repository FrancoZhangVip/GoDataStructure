package leetcode_algorithm

func LengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	checker := make(map[byte]int)
	input := []byte(s)
	p, q := 0, 1
	checker[input[p]] = p
	max := -1
	for q < len(input) && p < len(input) {
		if _, ok := checker[input[q]]; !ok {
			checker[input[q]] = q
			q++
			if q-p > max {
				max = q - p
			}
		} else {
			delete(checker, input[p])
			p++
		}

	}
	return max
}

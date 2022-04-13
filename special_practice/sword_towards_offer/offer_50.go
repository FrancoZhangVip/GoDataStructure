package sword_towards_offer

func FirstUniqChar(s string) byte {
	data := []byte(s)
	checker := make([]int, 26)
	for i := 0; i < len(data); i++ {
		checker[data[i]-'a']++
	}

	for i := 0; i < len(data); i++ {
		if checker[data[i]-'a'] == 1 {
			return data[i]
		}
	}
	return ' '
}

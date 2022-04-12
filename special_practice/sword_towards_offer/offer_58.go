package sword_towards_offer

func ReverseLeftWords(s string, n int) string {
	n %= len(s)
	k := len(s) - n
	data := []byte(s)

	for i := 0; i < len(data)/2; i++ {
		data[i], data[len(data)-1-i] = data[len(data)-1-i], data[i]
	}

	for i := 0; i < k/2; i++ {
		data[i], data[k-1-i] = data[k-1-i], data[i]
	}

	for i := k; i < (k+len(data))/2; i++ {
		data[i], data[len(data)-1-i+k] = data[len(data)-1-i+k], data[i]
	}

	return string(data)
}

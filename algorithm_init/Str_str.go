package algorithm_init

/*
strStr
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从 0 开始)。如果不存在，则返回 -1。
*/
func StrStr(haystack string, needle string) int {
	if len(haystack) == 0 {
		return -1
	}
	if len(needle) == 1 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}

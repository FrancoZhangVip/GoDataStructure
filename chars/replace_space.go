package chars

/*
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

示例 1：
输入：s = "We are happy."
输出："We%20are%20happy."
*/

func ReplaceSpace(s string) string {
	res := ""
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res = res + s[start:i] + "%20"
			start = i + 1
		}
	}
	res = res + s[start:]
	return res
}

/*
func replaceSpace(s string) string {
    b := []byte(s)
    result := make([]byte, 0)
    for i := 0; i < len(b); i++ {
        if b[i] == ' ' {
            result = append(result, []byte("%20")...)
        } else {
            result = append(result, b[i])
        }
    }
    return string(result)
}
*/

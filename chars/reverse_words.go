package chars

/*
给定一个字符串，逐个翻转字符串中的每个单词。

示例 1：
输入: "the sky is blue"
输出: "blue is sky the"

示例 2：
输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

示例 3：
输入: "a good   example"
输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
*/

func ReverseWords(s string) string {
	//1-reverse
	start := 0
	end := len(s) - 1
	data := []byte(s)
	for start < end {
		data[start], data[end] = data[end], data[start]
		start++
		end--
	}

	//2-delete muti space 、 head space and tail space
	headSpace := 0
	for i := 0; i < len(data); i++ {
		if data[i] == ' ' {
			headSpace++
		} else {
			break
		}
	}

	data = data[headSpace:]
	if len(data) <= 1 {
		return string(data)
	}

	start = 1
	for i := 1; i < len(data); i++ {
		if !(data[i] == ' ' && data[i-1] == ' ') {
			data[start] = data[i]
			start++
		}
	}

	if data[start-1] == ' ' {
		data = data[:start-1]
	} else {
		data = data[:start]
	}

	start = 0
	for i := 0; i < len(data); i++ {
		if data[i] == ' ' {
			end = i - 1
			for start < end {
				data[start], data[end] = data[end], data[start]
				start++
				end--
			}
			start = i + 1
		}
	}

	end = len(data) - 1
	for start < end {
		data[start], data[end] = data[end], data[start]
		start++
		end--
	}

	return string(data[:])
}

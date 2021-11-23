package chars

/*
给定一个字符串 s 和一个整数 k，你需要对从字符串开头算起的每隔 2k 个字符的前 k 个字符进行反转。
如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

示例:
输入: s = "abcdefg", k = 2
输出: "bacdfeg"
*/

func ReverseStr(s string, k int) string {
	data := []byte(s)

	for i := 0; i < len(data); i++ {
		modK := i % (2 * k)
		powK := i / k
		if modK < (k / 2) {
			smodK := i % k
			index := (powK+1)*k - 1
			if index > len(data)-1 {
				index = len(data) - 1
			}
			if index-smodK < i {
				continue
			}
			temp := data[index-smodK]
			data[index-smodK] = data[i]
			data[i] = temp
		}
	}
	return string(data)
}

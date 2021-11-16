package hash

/*
给你一个字符串数组 words ，请你找出所有在 words 的每个字符串中都出现的共用字符（ 包括重复字符），并以数组形式返回。你可以按 任意顺序 返回答案。

示例 1：
输入：words = ["bella","label","roller"] 输出：["e","l","l"]
示例 2：
输入：words = ["cool","lock","cook"] 输出：["c","o"]

提示：
1 <= words.length <= 100 1 <= words[i].length <= 100 words[i] 由小写英文字母组成
*/

func CommonChars(words []string) []string {
	if len(words) == 0 || len(words) == 1 {
		return words
	}
	set_size := 26
	res_set := make([][]int, len(words))
	for i := 0; i < len(words); i++ {
		res_set[i] = make([]int, set_size)
	}

	res := make([]string, 0)

	for j := 0; j < len(words); j++ {
		for i := 0; i < len(words[j]); i++ {
			set_index := int(words[j][i] - 'a')
			res_set[j][set_index]++
		}
	}

	for j := 0; j < set_size; j++ {
		min := len(words) * set_size
		for i := 0; i < len(words); i++ {
			if res_set[i][j] < min {
				min = res_set[i][j]
			}
		}
		if min != 0 {
			item := string(rune(int('a') + j))
			for k := 0; k < min; k++ {
				res = append(res, item)
			}

		}
	}

	return res
}

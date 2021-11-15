package hash

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

链接：https://leetcode-cn.com/problems/valid-anagram

示例 1:
输入: s = "anagram", t = "nagaram"
输出: true

示例 2:
输入: s = "rat", t = "car"
输出: false

*/

func IsAnagram(s string, t string) bool {
	set_size := 26
	sset := make([]int, set_size)
	tset := make([]int, set_size)

	for i := 0; i < len(s); i++ {
		set_index := int(s[i] - 'a')
		sset[set_index]++
	}

	for i := 0; i < len(t); i++ {
		set_index := int(t[i] - 'a')
		tset[set_index]++
	}

	for i := 0; i < set_size; i++ {
		if sset[i] != tset[i] {
			return false
		}
	}
	return true
}

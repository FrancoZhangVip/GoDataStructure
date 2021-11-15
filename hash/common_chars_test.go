package hash

import "testing"

func TestCommonChars(t *testing.T) {
	words := []string{"bella", "label", "roller"}
	res := CommonChars(words)
	ShowStrings(res)

}

func TestCommonChars2(t *testing.T) {
	words := []string{"cool", "lock", "cook"}
	res := CommonChars(words)
	ShowStrings(res)

}

func TestCommonChars3(t *testing.T) {
	words := []string{"acabcddd", "bcbdbcbd", "baddbadb", "cbdddcac", "aacbcccd", "ccccddda", "cababaab", "addcaccd"}
	res := CommonChars(words)
	ShowStrings(res)

}

func ShowStrings(strs []string) {
	for i := 0; i < len(strs); i++ {
		print(strs[i], "\t")
	}
	println()
}

func ShowArray(res [][]int) {
	for _, val := range res {
		for _, value := range val {
			print(value, "\t")
		}
		println()
	}
}

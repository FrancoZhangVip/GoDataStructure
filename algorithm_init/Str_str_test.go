package algorithm_init

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	hayStack := "gfsdfhjkwenlovehdjkgfsdgs,mnsjkdgj"
	needle := "love"

	fmt.Println(StrStr(hayStack, needle))
}

func TestStrStr_Not_Exist(t *testing.T) {
	hayStack := "gfsdfhjkwenlovehdjkgfsdgs,mnsjkdgj"
	needle := "loveyou"

	fmt.Println(StrStr(hayStack, needle))
}

func TestStrStr_AT_TOP(t *testing.T) {
	hayStack := "gfsdfhjkwenlovehdjkgfsdgs,mnsjkdgj"
	needle := "gfsdfh"

	fmt.Println(StrStr(hayStack, needle))
}

func TestStrStr_AT_END(t *testing.T) {
	hayStack := "gfsdfhjkwenlovehdjkgfsdgs,mnsjkdgj"
	needle := "kdgj"

	fmt.Println(len(hayStack))
	fmt.Println(StrStr(hayStack, needle))
}

func TestStrStr_AT_ALL(t *testing.T) {
	hayStack := "gfsdfhjkwenlovehdjkgfsdgs,mnsjkdgj"
	needle := hayStack

	fmt.Println(len(hayStack))
	fmt.Println(StrStr(hayStack, needle))
}

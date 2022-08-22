package hot_one_hundred

import (
	"fmt"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	res := FindDisappearedNumbers(nums)
	for _, val := range res {
		fmt.Print(val, "\t")
	}
}

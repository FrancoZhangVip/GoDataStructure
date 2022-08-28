package hot_one_hundred

import (
	"fmt"
	"testing"
)

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}
	ans := Subsets(nums)
	for _, data := range ans {
		fmt.Println(data)
	}

}

package sword_towards_offer

import (
	"fmt"
	"testing"
)

func TestFindRepeatNumber(t *testing.T) {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(FindRepeatNumber(nums))
}

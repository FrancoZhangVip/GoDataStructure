package sword_towards_offer

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	nums := []int{0, 1, 3}
	assert.Equal(t, 2, MissingNumber(nums))
}

func TestMissingNumber2(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	assert.Equal(t, 8, MissingNumber(nums))
}

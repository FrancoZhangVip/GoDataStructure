package sword_towards_offer

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	assert.Equal(t, 2, Search(nums, target))
}

func TestSearch2(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	assert.Equal(t, 0, Search(nums, target))
}

package array

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}

	assert.Equal(t, 2, MinSubArrayLen(target, nums))
}

package hot_one_hundred

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2}
	assert.Equal(t, 4, SingleNumber(nums))
}

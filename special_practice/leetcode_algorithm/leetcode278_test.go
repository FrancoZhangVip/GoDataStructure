package leetcode_algorithm

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	n := 5
	//bad := 4

	assert.Equal(t, 4, FirstBadVersion(n))
}

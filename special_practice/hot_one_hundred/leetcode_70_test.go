package hot_one_hundred

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	n := 2
	assert.Equal(t, 2, ClimbStairs(n))
}

func TestClimbStairs2(t *testing.T) {
	n := 3
	assert.Equal(t, 3, ClimbStairs(n))
}

func TestClimbStairs3(t *testing.T) {
	n := 44
	assert.Equal(t, 1134903170, ClimbStairs(n))
}

package sword_towards_offer

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestMinStackConstructor(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	res1 := minStack.Min()
	assert.Equal(t, -3, res1)
	minStack.Pop()
	res2 := minStack.Top()
	assert.Equal(t, 0, res2)
	res3 := minStack.Min()
	assert.Equal(t, -2, res3)
}

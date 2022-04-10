package sword_towards_offer

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	obj := CQueueConstructor()
	obj.AppendTail(3)
	res := obj.DeleteHead()

	assert.Equal(t, 3, res)

	res1 := obj.DeleteHead()

	assert.Equal(t, -1, res1)
}

func TestConstructor2(t *testing.T) {
	obj := CQueueConstructor()

	res1 := obj.DeleteHead()

	assert.Equal(t, -1, res1)

	obj.AppendTail(5)
	obj.AppendTail(2)

	res2 := obj.DeleteHead()

	assert.Equal(t, 5, res2)

	res3 := obj.DeleteHead()

	assert.Equal(t, 2, res3)
}

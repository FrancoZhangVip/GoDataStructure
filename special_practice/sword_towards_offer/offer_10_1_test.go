package sword_towards_offer

import (
	"fmt"
	"github.com/hedzr/assert"
	"testing"
)

func TestFib(t *testing.T) {
	assert.Equal(t, 1, Fib(2))
}

func TestFib2(t *testing.T) {
	assert.Equal(t, 5, Fib(5))
}

func TestFib3(t *testing.T) {
	fmt.Println(Fib(100))
}

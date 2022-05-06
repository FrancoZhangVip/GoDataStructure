package sword_towards_offer

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestMinArray(t *testing.T) {
	numbers := []int{3, 4, 5, 1, 2}
	assert.Equal(t, 1, MinArray(numbers))
}

func TestMinArray2(t *testing.T) {
	numbers := []int{2, 2, 2, 0, 1}
	assert.Equal(t, 0, MinArray(numbers))
}

func TestMinArray3(t *testing.T) {
	numbers := []int{2, 2, 2, 2, 2}
	assert.Equal(t, 2, MinArray(numbers))
}

func TestMinArray4(t *testing.T) {
	numbers := []int{2, 2, 0, 2, 2}
	assert.Equal(t, 0, MinArray(numbers))
}

func TestMinArray5(t *testing.T) {
	numbers := []int{1, 3, 5}
	assert.Equal(t, 1, MinArray(numbers))
}

func TestMinArray6(t *testing.T) {
	numbers := []int{1, 3, 3}
	assert.Equal(t, 1, MinArray(numbers))
}

func TestMinArray7(t *testing.T) {
	numbers := []int{3, 3, 1, 3}
	assert.Equal(t, 1, MinArray(numbers))
}

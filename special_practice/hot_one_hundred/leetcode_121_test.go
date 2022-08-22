package hot_one_hundred

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	data := []int{7, 1, 5, 3, 6, 4}
	profit := MaxProfit(data)
	assert.Equal(t, 5, profit)
}

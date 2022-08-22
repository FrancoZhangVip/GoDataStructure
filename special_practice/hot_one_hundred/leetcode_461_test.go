package hot_one_hundred

import (
	"fmt"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	x, y := 1, 4
	fmt.Println("distance is : ", HammingDistance(x, y))
}

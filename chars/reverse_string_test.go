package chars

import "testing"

func TestReverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	ReverseString(s)
}

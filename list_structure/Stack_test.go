package list_structure

import (
	"fmt"
	"testing"
)

func BenchmarkStack_Push(b *testing.B) {
	stack := NewStack()
	stack.Init()
	stack.Push(2)
	stack.Push(6)
	stack.Push(4)
	stack.Push(7)
	stack.Show()
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack()
	stack.Init()
	stack.Push(2)
	stack.Push(6)
	stack.Push(4)
	stack.Push(7)

	stack.Pop()
	stack.Show()
}

func TestStack_Size(t *testing.T) {
	stack := NewStack()
	stack.Init()
	stack.Push(2)
	fmt.Println(stack.Size())
}

package list_structure

import (
	"fmt"
	"testing"
)

func TestQueue_Push(t *testing.T) {
	queue := NewQueue()
	queue.Init()
	queue.Push(5)
	queue.Push(45)
	queue.Push(25)
	queue.Push(15)
	queue.Size()
	queue.Show()
}

func TestQueue_Pop(t *testing.T) {
	queue := NewQueue()
	queue.Init()
	queue.Push(4)
	e, err := queue.Pop()
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("element: ", e)
	}

	queue.Show()
}

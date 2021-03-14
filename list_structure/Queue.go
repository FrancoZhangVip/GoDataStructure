package list_structure

import (
	"errors"
	"fmt"
)

type Queue struct {
	queueElement []int
}

func NewQueue() *Queue {
	return new(Queue)
}

func (queue *Queue) Init() {
	queue.queueElement = make([]int, 0)
}

func (queue *Queue) Push(element int) {
	queue.queueElement = append(queue.queueElement, element)
}

func (queue *Queue) Pop() (element int, err error) {
	if len(queue.queueElement) == 0 {
		return 0, errors.New("empty error")
	}
	element = queue.queueElement[0]
	queue.queueElement = queue.queueElement[1:]
	return element, nil
}

func (queue *Queue) Size() (length int) {
	length = len(queue.queueElement)
	return length
}

func (queue *Queue) Show() {
	for _, e := range queue.queueElement {
		fmt.Print(e, "\t")
	}
	fmt.Println()
}

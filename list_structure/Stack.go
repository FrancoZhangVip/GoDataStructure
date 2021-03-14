package list_structure

import "fmt"

type Stack struct {
	stackElements []int
}

func NewStack() *Stack {
	return new(Stack)
}

func (stack *Stack) Init() {
	stack.stackElements = make([]int, 0)
}

func (stack *Stack) Push(element int) {
	stack.stackElements = append(stack.stackElements, element)
}

func (stack *Stack) Pop() (element int) {
	element = stack.stackElements[len(stack.stackElements)-1]
	stack.stackElements = stack.stackElements[0 : len(stack.stackElements)-1]
	return element
}

func (stack *Stack) Size() (length int) {
	length = len(stack.stackElements)
	return length
}

func (stack *Stack) Show() {
	for _, e := range stack.stackElements {
		fmt.Print(e, "\t")
	}
	fmt.Println()
}

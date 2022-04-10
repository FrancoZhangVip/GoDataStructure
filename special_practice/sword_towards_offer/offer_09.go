package sword_towards_offer

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

type CQueue struct {
	inStack, outStack     []int
	inputSize, outputSize int
}

func CQueueConstructor() CQueue {
	return CQueue{
		inStack:    make([]int, 0),
		outStack:   make([]int, 0),
		inputSize:  0,
		outputSize: 0,
	}
}

func (this *CQueue) AppendTail(value int) {
	if this.outputSize > 0 {
		this.inStack = make([]int, 0)
		this.inputSize = 0
		for i := this.outputSize - 1; i >= 0; i-- {
			this.inStack = append(this.inStack, this.outStack[i])
			this.inputSize++
		}

		this.outStack = make([]int, 0)
		this.outputSize = 0
	}
	this.inStack = append(this.inStack, value)
	this.inputSize++
}

func (this *CQueue) DeleteHead() int {
	if this.inputSize > 0 {
		this.outStack = make([]int, 0)
		this.outputSize = 0
		for i := this.inputSize - 1; i >= 0; i-- {
			this.outStack = append(this.outStack, this.inStack[i])
			this.outputSize++
		}
		this.inStack = make([]int, 0)
		this.inputSize = 0
	}
	if this.outputSize == 0 {
		return -1
	}
	res := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	this.outputSize--
	return res
}

package sword_towards_offer

import "math"

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

type MinStack struct {
	data    []int
	size    int
	min     int
	minData []int
	minSize int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:    make([]int, 0),
		size:    0,
		min:     math.MaxInt64,
		minData: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	if this.min > x {
		this.min = x
	}
	this.minData = append(this.minData, this.min)
	this.minSize++
	this.data = append(this.data, x)
	this.size++
}

func (this *MinStack) Pop() {

	this.minData = this.minData[:len(this.minData)-1]
	this.minSize--

	this.data = this.data[:len(this.data)-1]
	this.size--

	if len(this.minData) > 0 {
		this.min = this.minData[len(this.minData)-1]
	} else {
		this.min = math.MaxInt64
	}

}

func (this *MinStack) Top() int {
	return this.data[this.size-1]
}

func (this *MinStack) Min() int {
	return this.minData[this.minSize-1]
}

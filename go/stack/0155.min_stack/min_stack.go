package _155_min_stack

import "math"

type MinStack struct {
	stackData []int
	stackMin  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stackData: []int{},
		stackMin:  []int{math.MaxInt64},
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func (this *MinStack) Push(x int) {
	this.stackData = append(this.stackData, x)
	top := this.stackMin[len(this.stackMin)-1]
	this.stackMin = append(this.stackMin, min(top, x))
}

func (this *MinStack) Pop() {
	this.stackMin = this.stackMin[:len(this.stackMin)-1]
	this.stackData = this.stackData[:len(this.stackData)-1]
}

func (this *MinStack) Top() int {
	return this.stackData[len(this.stackData)-1]
}

func (this *MinStack) GetMin() int {
	return this.stackMin[len(this.stackMin)-1]
}

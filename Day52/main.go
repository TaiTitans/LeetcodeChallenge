package main

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		for len(this.inStack) > 0 {
			n := len(this.inStack)
			this.outStack = append(this.outStack, this.inStack[n-1])
			this.inStack = this.inStack[:n-1]
		}
	}
	n := len(this.outStack)
	res := this.outStack[n-1]
	this.outStack = this.outStack[:n-1]
	return res
}
func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		for len(this.inStack) > 0 {
			n := len(this.inStack)
			this.outStack = append(this.outStack, this.inStack[n-1])
			this.inStack = this.inStack[:n-1]
		}
	}
	n := len(this.outStack)
	return this.outStack[n-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

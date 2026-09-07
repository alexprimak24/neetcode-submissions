type MinStack struct {
	stack []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	stackLen := len(this.minStack)
	if stackLen == 0 || val <= this.minStack[stackLen-1] {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	lastEl := this.stack[len(this.stack) - 1]
	if lastEl == this.minStack[len(this.minStack) -1] {
		this.minStack = this.minStack[:len(this.minStack) - 1]
	}

	this.stack = this.stack[:len(this.stack) - 1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack) -1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack) -1]
}

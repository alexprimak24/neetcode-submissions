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
	this.stack = append(this.stack,val)

	if len(this.minStack) > 0 {
		if val <= this.minStack[len(this.minStack) - 1] {
			this.minStack = append(this.minStack, val)
		}
	} else {
		this.minStack = append(this.minStack, val)
		}
	
}

func (this *MinStack) Pop() {
	lastEl := this.stack[len(this.stack) - 1]

	this.stack = this.stack[:len(this.stack)-1]

	if lastEl == this.minStack[len(this.minStack) - 1] {
		this.minStack = this.minStack[:len(this.minStack) - 1]
	}
}

func (this *MinStack) Top() int {
	fmt.Println(len(this.minStack))
	return this.stack[len(this.stack) - 1]
}

func (this *MinStack) GetMin() int {
	fmt.Println(this.minStack)
	return this.minStack[len(this.minStack) - 1]
}

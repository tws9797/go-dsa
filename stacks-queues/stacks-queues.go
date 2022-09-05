package stacks_queues

type MyQueue struct {
	input  []int
	output []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.input = append(this.input, x)
}

func (this *MyQueue) Pop() int {
	tmp := this.Peek()
	this.output = this.output[:len(this.output)-1]
	return tmp
}

func (this *MyQueue) Peek() int {
	if len(this.output) == 0 {
		for len(this.input) > 0 {
			this.output = append(this.output, this.input[len(this.input)-1])
			this.input = this.input[:len(this.input)-1]
		}
	}

	return this.output[len(this.output)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.input) == 0 && len(this.output) == 0
}

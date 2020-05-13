package main

import "fmt"

type atom struct {
	min int
	val int
}

type MinStack struct {
	atoms []atom
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{atoms: []atom{}}
}

func (this *MinStack) top() atom {
	return this.atoms[len(this.atoms)-1]
}

func (this *MinStack) Push(x int) {
	var min int
	if len(this.atoms) == 0 {
		min = x
	} else {
		lastMin := this.top().min
		if x < lastMin {
			min = x
		} else {
			min = lastMin
		}
	}
	this.atoms = append(this.atoms, atom{min: min, val: x})
}

func (this *MinStack) Pop() {
	this.atoms = this.atoms[:len(this.atoms)-1]
}

func (this *MinStack) Top() int {
	return this.top().val
}

func (this *MinStack) GetMin() int {
	return this.top().min
}

func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	stack.Push(-1)
	stack.Push(10)
	fmt.Println(stack.Top())
	stack.Pop()
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.GetMin())
}

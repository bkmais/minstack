package main

import "fmt"

type MinStack struct {
	slice []int
	min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.slice) == 0 {
		this.min = append(this.min, val)
	} else {
		if val <= this.GetMin() {
			this.min = append(this.min, val)
		}
	}
	this.slice = append(this.slice, val)
}

func (this *MinStack) Pop() {
	if this.Top() == this.GetMin() {
		this.min = this.min[0 : len(this.min)-1]
	}
	this.slice = this.slice[0 : len(this.slice)-1]
}

func (this *MinStack) Top() int {
	return this.slice[len(this.slice)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) > 0 {
		return this.min[len(this.min)-1]
	} else {
		return 0
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {
	obj := Constructor()
	obj.Push(0)
	obj.Push(1)
	obj.Push(0)
	obj.Push(-2)
	obj.Push(4)
	//obj.Pop()
	param_1 := obj.Top()
	param_2 := obj.GetMin()

	fmt.Println("Top of the Stack:")
	fmt.Println(param_1)

	fmt.Println("Minimum in the Stack:")
	fmt.Println(param_2)

	fmt.Println(obj.slice)
	fmt.Println(obj.min)

}

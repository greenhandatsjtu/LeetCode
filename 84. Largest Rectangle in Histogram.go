package main

import "fmt"

// mono stack
func largestRectangleArea(heights []int) int {
	heights = append(heights, -1) //add bound
	monoStack := []int{-1}
	var square int
	for i, v := range heights {
		top := monoStack[len(monoStack)-1]
		for !(top == -1 || heights[top] < v) {
			left := monoStack[len(monoStack)-2]
			if cache := (i - left - 1) * heights[top]; cache > square {
				square = cache
			}
			monoStack = monoStack[:len(monoStack)-1] //pop
			top = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i) //push
	}
	return square
}

func main() {
	fmt.Println(largestRectangleArea([]int{1}))
}

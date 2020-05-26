package main

import (
	"fmt"
)

//Principle of Drawer
func findDuplicate(nums []int) int {
	low := 1
	n := len(nums)
	high := n - 1
	var mid, count int
	for low < high {
		count = 0
		mid = (high + low) / 2
		for _, v := range nums {
			if v <= mid {
				count++
			}
		}
		if count > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return (low + high) / 2
}

func main() {
	fmt.Println(findDuplicate([]int{2, 1, 3, 4, 2}))
}

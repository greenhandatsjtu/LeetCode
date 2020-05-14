package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}

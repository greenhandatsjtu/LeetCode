package main

import (
	"fmt"
	"sort"
)

//dynamic programming
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	copy(dp, nums)
	var count, cache int
	for i := 0; i < len(nums); i++ {
		count = 0
		for j := i - 2; j >= 0 && count < 2; j-- {
			cache = nums[i] + dp[j]
			if cache > dp[i] {
				dp[i] = cache
			}
			count++
		}
	}
	sort.Ints(dp)
	return dp[len(dp)-1]
}

func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

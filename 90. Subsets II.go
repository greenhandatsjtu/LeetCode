package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	subset := make([]int, 0)
	res := make([][]int, 0)
	sort.Ints(nums)
	backTrackDup(nums, 0, subset, &res)
	return res
}

func backTrackDup(nums []int, start int, subset []int, res *[][]int) {
	cache := make([]int, len(subset))
	copy(cache, subset)
	*res = append(*res, cache)
	for i := start; i < len(nums); i++ {
		if i != start && nums[i] == nums[i-1] {
			continue
		}
		subset = append(subset, nums[i])
		backTrackDup(nums, i+1, subset, res)
		subset = subset[:len(subset)-1]
	}
}

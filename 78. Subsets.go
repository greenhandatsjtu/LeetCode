package main

func subsets(nums []int) [][]int {
	subset := make([]int, 0)
	res := make([][]int, 0)
	backTrack(nums, 0, subset, &res)
	return res
}

func backTrack(nums []int, start int, subset []int, res *[][]int) {
	cache := make([]int, len(subset))
	copy(cache, subset)
	*res = append(*res, cache)
	for i := start; i < len(nums); i++ {
		subset = append(subset, nums[i])
		backTrack(nums, i+1, subset, res)
		subset = subset[:len(subset)-1]
	}
}

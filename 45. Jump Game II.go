package main

func jump(nums []int) int {
	var (
		rightmost int
		steps     int
		end       int
	)
	for i := 0; i < len(nums)-1; i++ {
		rightmost = max(rightmost, i+nums[i])
		if i == end {
			end = rightmost
			steps++
		}
	}
	return steps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

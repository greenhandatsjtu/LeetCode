package main

func canJump(nums []int) bool {
	rightmost := 0
	for i := range nums {
		if i <= rightmost {
			if i+nums[i] > rightmost {
				rightmost = i + nums[i]
			}
			if rightmost >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}

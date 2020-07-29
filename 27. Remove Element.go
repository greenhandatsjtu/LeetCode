package main

//double pointer
func removeElement(nums []int, val int) int {
	n := len(nums)
	slow := -1
	for fast := 0; fast < n; fast++ {
		if nums[fast] != val {
			slow += 1
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

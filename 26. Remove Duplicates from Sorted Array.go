package main

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	pre := nums[0]
	count := 0
	for i := 1; i < n; i++ {
		if pre == nums[i] {
			count++
		} else if count > 0 {
			n -= count
			for j := i - count; j < n; j++ {
				nums[j] = nums[j+count]
			}
			i -= count
			count = 0
		}
		pre = nums[i]
	}
	return n - count
}

func main() {
	removeDuplicates([]int{1, 1})
}

package main

//Backtracking algorithm
func permute(nums []int) [][]int {
	var (
		res          [][]int
		backTracking func()
	)
	n := len(nums)
	track := make([]int, 0, n)
	backTracking = func() {
		if len(track) == n {
			tmp := make([]int, n)
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for _, v := range nums {
			if isElement(track, v) {
				continue
			}
			track = append(track, v)
			backTracking()
			track = track[:len(track)-1]
		}
	}
	backTracking()
	return res
}

func isElement(track []int, t int) bool {
	for _, v := range track {
		if t == v {
			return true
		}
	}
	return false
}

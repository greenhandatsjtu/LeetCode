package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	odd := (m+n)%2 != 0
	mid := (m + n) / 2

	var i, j int
	index := -1
	var cache [2]int

	for i < m && j < n && index < mid {
		if nums1[i] < nums2[j] {
			cache[0], cache[1] = cache[1], nums1[i]
			i++
		} else {
			cache[0], cache[1] = cache[1], nums2[j]
			j++
		}
		index++
	}
	if i == m {
		for j < n && index < mid {
			cache[0], cache[1] = cache[1], nums2[j]
			j++
			index++
		}
	} else if j == n {
		for i < m && index < mid {
			cache[0], cache[1] = cache[1], nums1[i]
			i++
			index++
		}
	}

	if odd {
		return float64(cache[1])
	} else {
		return float64(cache[0]+cache[1]) / 2
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3}))
}

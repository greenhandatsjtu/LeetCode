package main

import (
	"fmt"
)

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left := 0
	right := -1
	start := 0
	end := -1
	count := 0
	window := make(map[byte]int)
	for {
		for right < len(s)-1 {
			right++
			c := s[right]
			window[c]++
			if window[c] == need[c] {
				count++
				if count == len(need) {
					break
				}
			}
		}
		if count == len(need) {
			for left <= right {
				c := s[left]
				window[c]--
				if window[c] < need[c] {
					if right-left < end-start || end == -1 {
						start, end = left, right
					}
					count--
					left++
					break
				}
				left++
			}
		}
		if right == len(s)-1 {
			break
		}
	}

	return s[start : end+1]
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "aa"))
}

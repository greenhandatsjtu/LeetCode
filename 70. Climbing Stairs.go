package main

func climbStairs(n int) int {
	path := make([]int, n+1)
	path[0], path[1] = 1, 1
	for i := 2; i <= n; i++ {
		path[i] = path[i-1] + path[i-2]
	}
	return path[n]
}

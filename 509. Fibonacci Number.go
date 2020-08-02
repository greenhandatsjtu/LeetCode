package main

//dynamic programming & state compress
func fib(N int) int {
	if N <= 0 {
		return 0
	}
	prev, curr := 1, 1
	for i := 3; i <= N; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}
	return curr
}

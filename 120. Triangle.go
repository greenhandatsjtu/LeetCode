package main

func minimumTotal(triangle [][]int) int {
	path := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if path[i] == nil {
				path[i] = make([]int, len(triangle[i]))
			}
			path[i][j] = triangle[i][j]
		}
	}
	for i := 1; i < len(path); i++ {
		for j := 0; j < len(path[i]); j++ {
			if j == 0 {
				path[i][j] += path[i-1][j]
			} else if j == len(path[i])-1 {
				path[i][j] += path[i-1][j-1]
			} else {
				path[i][j] += min(path[i-1][j-1], path[i-1][j])
			}
		}
	}
	result := path[len(path)-1][0]
	for i := 1; i < len(path[len(path)-1]); i++ {
		result = min(result, path[len(path)-1][i])
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

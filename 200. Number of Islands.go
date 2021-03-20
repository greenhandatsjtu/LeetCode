package main

func numIslands(grid [][]byte) int {
	count := 0
	var dfs func(int, int)
	dfs = func(i, j int) {
		grid[i][j] = '2'
		if i > 0 && grid[i-1][j] == '1' {
			dfs(i-1, j)
		}
		if j > 0 && grid[i][j-1] == '1' {
			dfs(i, j-1)
		}
		if i < len(grid)-1 && grid[i+1][j] == '1' {
			dfs(i+1, j)
		}
		if j < len(grid[0])-1 && grid[i][j+1] == '1' {
			dfs(i, j+1)
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}

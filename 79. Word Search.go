package main

type direction struct {
	x int
	y int
}

func exist(board [][]byte, word string) bool {
	directions := []direction{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	height, width := len(board), len(board[0])
	var search func(i, j, k int) bool
	search = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		cache := board[i][j]
		board[i][j] = '-'
		defer func() { board[i][j] = cache }()
		for _, dir := range directions {
			if x, y := i+dir.x, j+dir.y; 0 <= x && x < height && 0 <= y && y < width {
				if search(x, y, k+1) {
					return true
				}
			}
		}
		return false
	}

	for i, row := range board {
		for j := range row {
			if search(i, j, 0) {
				return true
			}
		}
	}
	return false
}

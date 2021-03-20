package main

func updateMatrix(matrix [][]int) [][]int {
	queue := make([][]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				matrix[i][j] = -1
			}
		}
	}
	dirs := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for len(queue) != 0 {
		point := queue[0]
		queue = queue[1:]
		for _, v := range dirs {
			x := point[0] + v[0]
			y := point[1] + v[1]
			if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) && matrix[x][y] == -1 {
				matrix[x][y] = matrix[point[0]][point[1]] + 1
				queue = append(queue, []int{x, y})
			}
		}
	}
	return matrix
}

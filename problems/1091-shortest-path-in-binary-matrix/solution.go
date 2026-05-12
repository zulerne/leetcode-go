// https://leetcode.com/problems/shortest-path-in-binary-matrix/description/
package shortestpathinbinarymatrix

func shortestPathBinaryMatrix(grid [][]int) int {
	res := 1
	n := len(grid)

	if grid[0][0] == 1 {
		return -1
	}
	if n == 1 {
		return 1
	}

	queue := [][2]int{{0, 0}}
	grid[0][0] = 1

	dirs := [8][2]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		{-1, -1}, {1, 1}, {1, -1}, {-1, 1},
	}

	for len(queue) > 0 {
		levelSize := len(queue)
		for range levelSize {
			i, j := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, d := range dirs {
				ni, nj := i+d[0], j+d[1]
				if ni >= 0 && nj >= 0 && ni < n && nj < n && grid[ni][nj] == 0 {
					if ni == n-1 && nj == n-1 {
						return res + 1
					}
					grid[ni][nj] = 1
					queue = append(queue, [2]int{ni, nj})
				}
			}
		}
		res++
	}

	return -1
}

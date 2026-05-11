// https://leetcode.com/problems/number-of-islands/description/
package numberofislands

func numIslands(grid [][]byte) int {
	var islands int
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				continue
			}

			islands++

			queue := [][2]int{{i, j}}
			for len(queue) > 0 {
				ii, jj := queue[0][0], queue[0][1]
				queue = queue[1:]
				grid[ii][jj] = '0'

				for _, d := range dirs {
					ni, nj := ii+d[0], jj+d[1]
					if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[ni]) && grid[ni][nj] == '1' {
						grid[ni][nj] = '0'
						queue = append(queue, [2]int{ni, nj})
					}
				}
			}
		}
	}

	return islands
}

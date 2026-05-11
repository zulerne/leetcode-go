// https://leetcode.com/problems/number-of-islands/description/
package numberofislands

func numIslands(grid [][]byte) int {
	var res int
	n, m := len(grid), len(grid[0])
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				continue
			}
			grid[i][j] = '0'
			res++

			queue := [][2]int{{i, j}}
			for len(queue) > 0 {
				first := queue[0]
				queue = queue[1:]
				ii, jj := first[0], first[1]

				for _, d := range dirs {
					ni, nj := ii+d[0], jj+d[1]
					if ni >= 0 && ni < n && nj >= 0 && nj < m && grid[ni][nj] == '1' {
						grid[ni][nj] = '0'
						queue = append(queue, [2]int{ni, nj})
					}
				}
			}
		}
	}

	return res
}

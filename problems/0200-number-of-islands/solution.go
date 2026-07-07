// https://leetcode.com/problems/number-of-islands/description/
package numberofislands

func numIslands(grid [][]byte) int {
	var res int
	n, m := len(grid), len(grid[0])
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for i := range n {
		for j := range m {
			if grid[i][j] == '0' {
				continue
			}

			res++
			grid[i][j] = '0'
			queue := [][2]int{{i, j}}

			for len(queue) > 0 {
				curI, curJ := queue[0][0], queue[0][1]
				queue = queue[1:]

				for _, dir := range dirs {
					ni, nj := curI+dir[0], curJ+dir[1]

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

func numIslandsDfs(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	n, m := len(grid), len(grid[0])
	res := 0

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] == '0' {
			return
		}

		grid[i][j] = '0'

		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	for i := range n {
		for j := range m {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}

	return res
}

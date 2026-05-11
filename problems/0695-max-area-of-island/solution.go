// https://leetcode.com/problems/max-area-of-island/description/
package maxareaofisland

func maxAreaOfIsland(grid [][]int) int {
	var maxArea int
	n, m := len(grid), len(grid[0])
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for i := range n {
		for j := range m {
			if grid[i][j] == 0 {
				continue
			}
			queue := [][2]int{{i, j}}
			grid[i][j] = 0
			var curArea int
			for len(queue) > 0 {
				ii, jj := queue[0][0], queue[0][1]
				queue = queue[1:]

				curArea++

				for _, d := range dirs {
					ni, nj := ii+d[0], jj+d[1]
					if ni >= 0 && nj >= 0 && ni < n && nj < m && grid[ni][nj] == 1 {
						grid[ni][nj] = 0
						queue = append(queue, [2]int{ni, nj})
					}
				}
			}
			maxArea = max(maxArea, curArea)
		}
	}

	return maxArea
}

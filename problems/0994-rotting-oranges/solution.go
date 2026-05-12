// https://leetcode.com/problems/rotting-oranges/description/
package rottingoranges

func orangesRotting(grid [][]int) int {
	var res int
	var fresh int
	n, m := len(grid), len(grid[0])
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	queue := [][2]int{}
	for i := range n {
		for j := range m {
			switch grid[i][j] {
			case 2:
				queue = append(queue, [2]int{i, j})
			case 1:
				fresh++
			}
		}
	}

	for len(queue) > 0 {
		wasRotten := false
		levelSize := len(queue)
		for range levelSize {
			i, j := queue[0][0], queue[0][1]
			queue = queue[1:]

			for _, d := range dirs {
				ni, nj := i+d[0], j+d[1]

				if ni < 0 || nj < 0 || ni >= n || nj >= m || grid[ni][nj] != 1 {
					continue
				}

				wasRotten = true
				fresh--
				grid[ni][nj] = 2
				queue = append(queue, [2]int{ni, nj})
			}
		}
		if wasRotten {
			res++
		}
	}

	if fresh > 0 {
		return -1
	}
	return res
}

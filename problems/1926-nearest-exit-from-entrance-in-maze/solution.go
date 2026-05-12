// https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/description/
package nearestexitfromentranceinmaze

func nearestExit(maze [][]byte, entrance []int) int {
	var res int
	n, m := len(maze), len(maze[0])
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	queue := [][2]int{{entrance[0], entrance[1]}}
	maze[entrance[0]][entrance[1]] = '+'
	for len(queue) > 0 {
		levelSize := len(queue)
		for range levelSize {
			i, j := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, d := range dirs {
				ni, nj := i+d[0], j+d[1]

				if ni >= 0 && nj >= 0 && ni < n && nj < m && maze[ni][nj] == '.' {
					if ni == 0 || ni == n-1 || nj == 0 || nj == m-1 {
						return res + 1
					}

					maze[ni][nj] = '+'
					queue = append(queue, [2]int{ni, nj})
				}

			}
		}
		res++
	}

	return -1
}

// https://leetcode.com/problems/01-matrix/description/
package onematrix

import "math"

// TODO: Repeat
func updateMatrix(mat [][]int) [][]int {
	n, m := len(mat), len(mat[0])
	res := make([][]int, n)
	for i := range n {
		res[i] = make([]int, m)
	}
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	queue := [][2]int{}
	for i := range n {
		for j := range m {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
				res[i][j] = 0
			} else {
				res[i][j] = math.MaxInt
			}
		}
	}

	for len(queue) > 0 {
		i, j := queue[0][0], queue[0][1]
		queue = queue[1:]

		for _, d := range dirs {
			ni, nj := i+d[0], j+d[1]

			if ni < 0 || nj < 0 || ni >= n || nj >= m || res[ni][nj] <= res[i][j]+1 {
				continue
			}
			res[ni][nj] = res[i][j] + 1
			queue = append(queue, [2]int{ni, nj})
		}
	}

	return res
}

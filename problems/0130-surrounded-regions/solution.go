// https://leetcode.com/problems/surrounded-regions/description/
package surroundedregions

func solve(board [][]byte) {
	n, m := len(board), len(board[0])
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	queue := [][2]int{}

	for i := range n {
		if board[i][0] == 'O' {
			board[i][0] = 'T'
			queue = append(queue, [2]int{i, 0})
		}
		if board[i][m-1] == 'O' {
			board[i][m-1] = 'T'
			queue = append(queue, [2]int{i, m - 1})
		}
	}

	for j := range m {
		if board[0][j] == 'O' {
			board[0][j] = 'T'
			queue = append(queue, [2]int{0, j})
		}
		if board[n-1][j] == 'O' {
			board[n-1][j] = 'T'
			queue = append(queue, [2]int{n - 1, j})
		}
	}

	for len(queue) > 0 {
		i, j := queue[0][0], queue[0][1]
		queue = queue[1:]

		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]

			if ni >= 0 && nj >= 0 && ni < n && nj < m && board[ni][nj] == 'O' {
				board[ni][nj] = 'T'
				queue = append(queue, [2]int{ni, nj})
			}
		}
	}

	for i := range n {
		for j := range m {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X'
			case 'T':
				board[i][j] = 'O'
			}
		}
	}
}

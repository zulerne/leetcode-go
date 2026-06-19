// https://leetcode.com/problems/game-of-life/description/
package gameoflife

func gameOfLife(board [][]int) {
	n, m := len(board), len(board[0])

	dirs := [8][2]int{
		[2]int{-1, -1},
		[2]int{-1, 0},
		[2]int{-1, 1},

		[2]int{0, -1},
		[2]int{0, 1},

		[2]int{1, -1},
		[2]int{1, 0},
		[2]int{1, 1},
	}

	for i := range n {
		for j := range m {
			liveN := 0

			for _, d := range dirs {
				ni, nj := i+d[0], j+d[1]
				if ni >= 0 && ni < n && nj >= 0 && nj < m {
					if board[ni][nj] == 1 || board[ni][nj] == 2 {
						liveN++
					}
				}
			}

			if board[i][j] == 1 && (liveN < 2 || liveN > 3) {
				board[i][j] = 2
			} else if board[i][j] == 0 && liveN == 3 {
				board[i][j] = -1
			}
		}
	}

	for i := range n {
		for j := range m {
			switch board[i][j] {
			case 2:
				board[i][j] = 0
			case -1:
				board[i][j] = 1
			}
		}
	}
}

// https://leetcode.com/problems/n-queens/description/
package nqueens

import (
	"slices"
)

func solveNQueens(n int) [][]string {
	var res [][]string

	joinBoard := func(board [][]byte) []string {
		res := make([]string, n)

		for i := range n {
			acc := make([]byte, n)
			for j := range n {
				acc[j] = board[i][j]
			}
			res[i] = string(acc)
		}

		return res
	}

	var backtrack func(row int, columns, diagonals, tDiagonals []int, board [][]byte)
	backtrack = func(row int, columns, diagonals, tDiagonals []int, board [][]byte) {
		if row == n {
			res = append(res, joinBoard(board))
			return
		}

		for col := range n {
			curDiagonal := row + col
			curTDiagonal := row - col

			if slices.Contains(columns, col) || slices.Contains(diagonals, curDiagonal) || slices.Contains(tDiagonals, curTDiagonal) {
				continue
			}

			columns = append(columns, col)
			diagonals = append(diagonals, curDiagonal)
			tDiagonals = append(tDiagonals, curTDiagonal)
			board[row][col] = 'Q'

			backtrack(row+1, columns, diagonals, tDiagonals, board)

			columns = columns[:len(columns)-1]
			diagonals = diagonals[:len(diagonals)-1]
			tDiagonals = tDiagonals[:len(tDiagonals)-1]
			board[row][col] = '.'
		}
	}

	board := make([][]byte, n)
	for i := range n {
		board[i] = make([]byte, n)
		for j := range n {
			board[i][j] = '.'
		}
	}

	backtrack(0, []int{}, []int{}, []int{}, board)

	return res
}

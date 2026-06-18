// https://leetcode.com/problems/valid-sudoku/description/
package validsudoku

func isValidSudoku(board [][]byte) bool {

	var rows [9][10]bool
	var cols [9][10]bool
	var boxes [9][10]bool

	for r := range 9 {
		for c := range 9 {
			if board[r][c] == '.' {
				continue
			}

			num := board[r][c] - '0'
			boardIdx := (r/3)*3 + (c / 3)
			if rows[r][num] || cols[c][num] || boxes[boardIdx][num] {
				return false
			}

			rows[r][num] = true
			cols[c][num] = true
			boxes[boardIdx][num] = true
		}
	}

	return true
}

// https://leetcode.com/problems/word-search/description/
package wordsearch

func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var backtrack func(i, j, pos int) bool
	backtrack = func(i, j, pos int) bool {
		if pos == len(word)-1 {
			return board[i][j] == word[pos]
		}
		if board[i][j] != word[pos] {
			return false
		}

		old := board[i][j]
		board[i][j] = '#'
		defer func() { board[i][j] = old }()
		for _, d := range dirs {
			ni, nj := i+d[0], j+d[1]

			if ni >= 0 && ni < n && nj >= 0 && nj < m {
				if backtrack(ni, nj, pos+1) {
					return true
				}
			}
		}

		return false
	}

	for i := range n {
		for j := range m {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}

	return false
}

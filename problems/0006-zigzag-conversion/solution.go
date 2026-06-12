// https://leetcode.com/problems/zigzag-conversion/description/
package zigzagconversion

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	cell := make([][]byte, numRows)
	row, up := 0, true

	for i := range s {
		cell[row] = append(cell[row], s[i])

		if row == 0 || row == numRows-1 {
			up = !up
		}

		if up {
			row -= 1
		} else {
			row += 1
		}
	}

	var res strings.Builder
	res.Grow(len(s))
	for _, v := range cell {
		res.Write(v)
	}
	return res.String()
}

// https://leetcode.com/problems/set-matrix-zeroes/description/
package setmatrixzeroes

func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])

	iZ := make([]bool, n)
	jZ := make([]bool, m)

	for i := range n {
		for j := range m {
			if matrix[i][j] == 0 {
				iZ[i] = true
				jZ[j] = true
			}
		}
	}

	for i := range n {
		for j := range m {
			if iZ[i] || jZ[j] {
				matrix[i][j] = 0
			}
		}
	}
}

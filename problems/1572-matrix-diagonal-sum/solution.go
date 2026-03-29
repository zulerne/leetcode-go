// https://leetcode.com/problems/matrix-diagonal-sum/description/
package matrixdiagonalsum

/*
func diagonalSum(mat [][]int) int {
	var sum int
	len := len(mat)
	for i := range len {
		for j := range len {
			if i == j || i == len-1-j {
				sum += mat[i][j]
			}
		}
	}
	return sum
}
*/

func diagonalSum(mat [][]int) int {
	var sum int
	len := len(mat)

	for i := range len {
		sum += mat[i][i] + mat[i][len-1-i]
	}
	if len%2 == 1 {
		sum -= mat[len/2][len/2]
	}
	return sum
}

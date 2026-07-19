// https://leetcode.com/problems/search-a-2d-matrix/description/
package searcha2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	left, right := 0, n*m-1

	for left <= right {
		mid := left + (right-left)/2
		row, col := mid/m, mid%m

		if matrix[row][col] == target {
			return true
		}

		if matrix[row][col] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

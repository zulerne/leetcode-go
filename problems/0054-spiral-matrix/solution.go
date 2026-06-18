package spiralmatrix

func spiralOrder(matrix [][]int) []int {
	rows, cols := len(matrix), len(matrix[0])
	res := make([]int, 0, rows*cols)
	top, bottom := 0, rows-1
	left, right := 0, cols-1

	for len(res) < rows*cols {
		// right
		for j := left; j <= right; j++ {
			res = append(res, matrix[top][j])
		}
		top++

		// down
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		if top > bottom || left > right {
			break
		}

		// left
		for j := right; j >= left; j-- {
			res = append(res, matrix[bottom][j])
		}
		bottom--

		// up
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}

	return res
}

// https://leetcode.com/problems/snakes-and-ladders/description
package snakesandladders

import (
	"math"
)

func snakesAndLadders(board [][]int) int {
	var res int
	n := len(board)

	numToCeil := func(num int) (int, int) {
		var i, j int

		i = n - int(math.Ceil(float64(num)/float64(n)))
		reverse := (n-1-i)%2 != 0

		if reverse {
			if num%n == 0 {
				j = 0
			} else {
				j = n - num%n
			}
		} else {
			if num%n == 0 {
				j = n - 1
			} else {
				j = num%n - 1
			}
		}

		return i, j
	}

	seen := make([]bool, n*n+1)
	seen[1] = true

	queue := []int{1}
	for len(queue) > 0 {
		lvlSize := len(queue)

		for range lvlSize {
			cur := queue[0]
			queue = queue[1:]

			if cur == n*n {
				return res
			}

			for dice := cur + 1; dice <= cur+6 && dice <= n*n; dice++ {
				nextI, nextJ := numToCeil(dice)
				dest := dice
				if board[nextI][nextJ] != -1 {
					dest = board[nextI][nextJ]
				}

				if seen[dest] {
					continue
				}

				seen[dest] = true

				queue = append(queue, dest)
			}
		}

		res++
	}

	return -1
}

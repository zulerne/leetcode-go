// https://leetcode.com/problems/climbing-stairs/description/
package climbingstairs

func climbStairs(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for range n {
		a, b = b, a+b
	}

	return b
}

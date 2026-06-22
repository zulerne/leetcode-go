// https://leetcode.com/problems/happy-number/description/
package happynumber

func isHappy(n int) bool {
	visited := make(map[int]bool)

	for n != 1 {
		if visited[n] {
			return false
		}

		visited[n] = true

		sum := 0
		for n != 0 {
			d := n % 10
			sum += d * d
			n /= 10
		}

		n = sum
	}

	return true
}

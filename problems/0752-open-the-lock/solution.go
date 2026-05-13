// https://leetcode.com/problems/open-the-lock/description/
package openthelock

func openLock(deadends []string, target string) int {
	var res int
	queue := []string{"0000"}
	seen := make(map[string]bool, len(deadends))

	neighs := make([]string, 8)
	updateNeighs := func(cur string) {
		b := []byte(cur)
		m := len(neighs) / 2

		for i := range m {
			num := int(cur[i] - '0')
			x := (num + 1) % 10
			b[i] = byte(x + '0')
			neighs[i] = string(b)
			x = (10 + num - 1) % 10
			b[i] = byte(x + '0')
			neighs[m+i] = string(b)
			b[i] = cur[i]
		}
	}

	for _, d := range deadends {
		seen[d] = true
	}
	if seen["0000"] {
		return -1
	}

	for len(queue) > 0 {
		levelSize := len(queue)

		for range levelSize {
			cur := queue[0]
			queue = queue[1:]

			if cur == target {
				return res
			}
			updateNeighs(cur)
			for _, n := range neighs {
				if !seen[n] {
					seen[n] = true
					queue = append(queue, n)
				}
			}
		}
		res++
	}

	return -1
}

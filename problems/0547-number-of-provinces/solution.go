// https://leetcode.com/problems/number-of-provinces/description/
package numberofprovinces

func findCircleNum(isConnected [][]int) int {
	var provinces int
	n := len(isConnected)
	seen := make([]bool, n)

	for city := range n {
		if seen[city] {
			continue
		}

		seen[city] = true
		provinces++

		queue := []int{city}
		for len(queue) > 0 {
			first := queue[0]
			queue = queue[1:]

			for neigh, isConn := range isConnected[first] {
				if !seen[neigh] && isConn == 1 {
					seen[neigh] = true
					queue = append(queue, neigh)
				}
			}
		}
	}

	return provinces
}

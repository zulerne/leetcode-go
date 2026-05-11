// https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/description/
package numberofconnectedcomponentsinanundirectedgraph

func CountComponents(n int, edges [][]int) int {
	var res int

	graph := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	seen := make([]bool, n)

	for node := range n {
		if seen[node] {
			continue
		}

		seen[node] = true
		res++

		stack := []int{node}
		for len(stack) > 0 {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			for _, neigh := range graph[last] {
				if !seen[neigh] {
					seen[neigh] = true
					stack = append(stack, neigh)
				}
			}
		}
	}

	return res
}

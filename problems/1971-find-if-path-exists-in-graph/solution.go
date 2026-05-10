// https://leetcode.com/problems/find-if-path-exists-in-graph/description/
package findifpathexistsingraph

func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := make([][]int, n)
	stack := []int{source}
	seen := make([]bool, n)

	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	seen[source] = true

	for len(stack) > 0 {
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if vertex == destination {
			return true
		}

		for _, neighbor := range graph[vertex] {
			if !seen[neighbor] {
				seen[neighbor] = true
				stack = append(stack, neighbor)
			}
		}
	}

	return false
}

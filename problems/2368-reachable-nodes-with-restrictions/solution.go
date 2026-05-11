// https://leetcode.com/problems/reachable-nodes-with-restrictions/description/
package reachablenodeswithrestrictions

func reachableNodes(n int, edges [][]int, restricted []int) int {
	var res int
	graph := make([][]int, n)
	visited := make([]bool, n)

	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	for _, r := range restricted {
		visited[r] = true
	}

	queue := []int{0}
	visited[0] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res++

		for _, neigh := range graph[node] {
			if !visited[neigh] {
				visited[neigh] = true
				queue = append(queue, neigh)
			}
		}
	}

	return res
}

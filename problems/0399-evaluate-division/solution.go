// https://leetcode.com/problems/evaluate-division/description/
package evaluatedivision

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	res := make([]float64, len(queries))

	for i := range res {
		res[i] = -1
	}

	graph := make(map[string]map[string]float64)

	for i, eq := range equations {
		a, b := eq[0], eq[1]

		if graph[a] == nil {
			graph[a] = make(map[string]float64)
		}
		if graph[b] == nil {
			graph[b] = make(map[string]float64)
		}
		graph[a][b] = values[i]
		graph[b][a] = 1 / values[i]
	}

	type step struct {
		node string
		val  float64
	}
	queue := []step{}
	seen := make(map[string]bool)

	for i, q := range queries {
		a, b := q[0], q[1]

		if graph[a] == nil || graph[b] == nil {
			continue
		}

		queue = queue[:0]
		clear(seen)

		queue = append(queue, step{node: a, val: 1})
		seen[a] = true

		for len(queue) > 0 {
			node, val := queue[0].node, queue[0].val
			queue = queue[1:]

			if node == b {
				res[i] = val
				break
			}

			for nextNode, weight := range graph[node] {
				if !seen[nextNode] {
					queue = append(queue, step{node: nextNode, val: val * weight})
					seen[nextNode] = true
				}
			}
		}
	}

	return res
}

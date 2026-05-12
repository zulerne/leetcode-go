// https://leetcode.com/problems/shortest-path-with-alternating-colors/description/
package shortestpathwithalternatingcolors

// TODO: Repeat
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	res := make([]int, n)
	graph := make(map[bool]map[int][]int)

	graph[true] = make(map[int][]int, len(redEdges))
	graph[false] = make(map[int][]int, len(blueEdges))
	for _, edge := range redEdges {
		graph[true][edge[0]] = append(graph[true][edge[0]], edge[1])
	}
	for _, edge := range blueEdges {
		graph[false][edge[0]] = append(graph[false][edge[0]], edge[1])
	}

	type State struct {
		node  int
		isRed bool
	}
	queue := []State{{0, true}, {0, false}}
	seen := make(map[State]bool)
	count := 0
	for len(queue) > 0 {
		levelSize := len(queue)
		count++

		for range levelSize {
			node, flag := queue[0].node, queue[0].isRed
			queue = queue[1:]

			neighbs := graph[flag][node]
			for _, n := range neighbs {
				if res[n] <= 0 {
					res[n] = count
				}
				state := State{node: n, isRed: !flag}
				if !seen[state] {
					seen[state] = true
					queue = append(queue, state)
				}
			}
		}
	}

	res[0] = 0
	for i := 1; i < len(res); i++ {
		if res[i] == 0 {
			res[i] = -1
		}
	}

	return res
}

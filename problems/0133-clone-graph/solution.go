// https://leetcode.com/problems/clone-graph/description/
package clonegraph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	graph := make(map[*Node]*Node)
	graph[node] = &Node{Val: node.Val}
	queue := []*Node{node}

	for len(queue) > 0 {
		oldN := queue[0]
		queue = queue[1:]

		for _, neigh := range oldN.Neighbors {
			if _, ok := graph[neigh]; !ok {
				graph[neigh] = &Node{Val: neigh.Val}
				queue = append(queue, neigh)
			}
			graph[oldN].Neighbors = append(graph[oldN].Neighbors, graph[neigh])
		}
	}

	return graph[node]
}

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

	nodes := make(map[*Node]*Node)
	queue := []*Node{node}

	nodes[node] = &Node{Val: node.Val}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		for _, neigh := range n.Neighbors {
			_, ok := nodes[neigh]
			if !ok {
				nodes[neigh] = &Node{Val: neigh.Val}
				queue = append(queue, neigh)
			}

			nodes[n].Neighbors = append(nodes[n].Neighbors, nodes[neigh])
		}
	}

	return nodes[node]
}

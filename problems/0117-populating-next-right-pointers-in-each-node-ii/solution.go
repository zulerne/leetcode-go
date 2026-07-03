// https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/description/
package populatingnextrightpointersineachnodeii

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		lvlSize := len(queue)

		for i := range lvlSize {
			node := queue[0]
			queue = queue[1:]

			if i < lvlSize-1 {
				node.Next = queue[0]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}

func connectMemory(root *Node) *Node {
	if root == nil {
		return nil
	}

	parent := root

	for parent != nil {
		dummy := &Node{}
		tail := dummy

		for parent != nil {
			if parent.Left != nil {
				tail.Next = parent.Left
				tail = tail.Next
			}
			if parent.Right != nil {
				tail.Next = parent.Right
				tail = tail.Next
			}
			parent = parent.Next
		}

		parent = dummy.Next
	}

	return root
}

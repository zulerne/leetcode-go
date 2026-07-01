// https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
package maximumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type nodeDepth struct {
	node  *TreeNode
	depth int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var res int
	stack := []nodeDepth{{node: root, depth: 1}}

	for len(stack) > 0 {
		nd := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node, depth := nd.node, nd.depth

		res = max(res, depth)

		if node.Left != nil {
			stack = append(stack, nodeDepth{node: node.Left, depth: depth + 1})
		}
		if node.Right != nil {
			stack = append(stack, nodeDepth{node: node.Right, depth: depth + 1})
		}
	}

	return res
}

func maxDepthRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

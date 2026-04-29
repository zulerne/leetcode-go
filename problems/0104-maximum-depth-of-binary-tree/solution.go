// https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
package maximumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

type stackElem struct {
	node  *TreeNode
	depth int
}

func maxDepthIterative(root *TreeNode) int {
	var result int
	stack := []stackElem{{node: root, depth: 1}}
	for len(stack) > 0 {
		elem := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node, depth := elem.node, elem.depth
		if node == nil {
			continue
		}
		stack = append(stack, stackElem{node: node.Left, depth: depth + 1})
		stack = append(stack, stackElem{node: node.Right, depth: depth + 1})
		result = max(result, depth)
	}
	return result
}

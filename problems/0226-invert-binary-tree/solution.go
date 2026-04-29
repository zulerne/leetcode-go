// https://leetcode.com/problems/invert-binary-tree/description/
package invertbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTreeIterative(root *TreeNode) *TreeNode {
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node == nil {
			continue
		}

		node.Left, node.Right = node.Right, node.Left
		stack = append(stack, node.Left)
		stack = append(stack, node.Right)
	}
	return root
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

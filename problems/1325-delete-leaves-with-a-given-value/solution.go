// https://leetcode.com/problems/delete-leaves-with-a-given-value/description/
package deleteleaveswithagivenvalue

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}

	return root
}

func removeLeafNodesIterative(root *TreeNode, target int) *TreeNode {
	dummy := &TreeNode{Left: root}
	stack := []*TreeNode{dummy}
	order := []*TreeNode{}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		order = append(order, node)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	for i := len(order) - 1; i >= 0; i-- {
		node := order[i]
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil && node.Left.Val == target {
			node.Left = nil
		}
		if node.Right != nil && node.Right.Left == nil && node.Right.Right == nil && node.Right.Val == target {
			node.Right = nil
		}
	}

	return dummy.Left
}

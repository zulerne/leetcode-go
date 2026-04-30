// https://leetcode.com/problems/balanced-binary-tree/description/
package balancedbinarytree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var height func(root *TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return max(height(root.Left), height(root.Right)) + 1
	}

	if root == nil {
		return true
	}
	if int(math.Abs(float64(height(root.Left)-height(root.Right)))) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

// TODO: add isBalancedIterative

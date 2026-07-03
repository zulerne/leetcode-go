// https://leetcode.com/problems/count-complete-tree-nodes/description/
package countcompletetreenodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := 0
	left := root
	for left != nil {
		leftHeight++
		left = left.Left
	}

	rightHeight := 0
	right := root
	for right != nil {
		rightHeight++
		right = right.Right
	}

	if leftHeight == rightHeight {
		return (1 << leftHeight) - 1
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}

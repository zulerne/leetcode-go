// https://leetcode.com/problems/validate-binary-search-tree/description/
package validatebinarysearchtree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	cur := math.MinInt
	var helper func(root *TreeNode) bool
	helper = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if !helper(root.Left) {
			return false
		}
		if root.Val <= cur {
			return false
		}
		cur = root.Val
		return helper(root.Right)
	}

	return helper(root)
}

type stackElem struct {
	node       *TreeNode
	leftBound  int
	rightBound int
}

func isValidBSTIterative(root *TreeNode) bool {
	stack := []stackElem{{node: root, leftBound: math.MinInt, rightBound: math.MaxInt}}

	for len(stack) > 0 {
		el := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node, leftBound, rightBound := el.node, el.leftBound, el.rightBound

		if node == nil {
			continue
		}
		if node.Val <= leftBound || node.Val >= rightBound {
			return false
		}
		stack = append(stack, stackElem{node: node.Left, leftBound: leftBound, rightBound: node.Val})
		stack = append(stack, stackElem{node: node.Right, leftBound: node.Val, rightBound: rightBound})
	}

	return true
}

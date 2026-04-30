// https://leetcode.com/problems/insert-into-a-binary-search-tree/description/
package insertintoabinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

func insertIntoBSTIterative(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	if root == nil {
		return node
	}
	cur := root
	for {
		if val < cur.Val {
			if cur.Left == nil {
				cur.Left = node
				return root
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = node
				return root
			}
			cur = cur.Right
		}
	}
}

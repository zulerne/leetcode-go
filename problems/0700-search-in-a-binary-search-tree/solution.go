// https://leetcode.com/problems/search-in-a-binary-search-tree/description/
package searchinabinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	} else if val > root.Val {
		return searchBST(root.Right, val)
	}
	return root
}

func searchBSTIterative(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	cur := root
	for cur != nil {
		if val < cur.Val {
			cur = cur.Left
		} else if val > cur.Val {
			cur = cur.Right
		} else {
			return cur
		}
	}
	return cur
}

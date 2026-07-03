// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/description/
package flattenbinarytreetolinkedlist

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	cur := root

	for cur != nil {
		if cur.Left != nil {
			rightmost := cur.Left
			for rightmost.Right != nil {
				rightmost = rightmost.Right
			}

			rightmost.Right = cur.Right

			cur.Right = cur.Left
			cur.Left = nil
		}

		cur = cur.Right
	}
}

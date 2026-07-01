// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
package constructbinarytreefrompreorderandinordertraversal

import (
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	rootI := slices.Index(inorder, root.Val)

	root.Left = buildTree(preorder[1:2], inorder[:rootI])
	root.Right = buildTree(preorder[2:], inorder[rootI+1:])

	return root
}

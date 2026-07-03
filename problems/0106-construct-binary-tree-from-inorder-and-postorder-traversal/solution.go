// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
package constructbinarytreefrominorderandpostordertraversal

import (
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	rootI := slices.Index(inorder, root.Val)

	root.Left = buildTree(inorder[:rootI], postorder[:rootI])
	root.Right = buildTree(inorder[rootI+1:], postorder[rootI:len(postorder)-1])

	return root
}

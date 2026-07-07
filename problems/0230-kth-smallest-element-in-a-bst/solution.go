// https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/
package kthsmallestelementinabst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var res int
	i := 0
	found := false

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil || found {
			return
		}

		inorder(root.Left)

		i++
		if i == k {
			res = root.Val
			found = true
			return
		}

		inorder(root.Right)
	}

	inorder(root)

	return res
}

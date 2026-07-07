// https://leetcode.com/problems/minimum-absolute-difference-in-bst/description/
package minimumabsolutedifferenceinbst

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	res := math.MaxInt

	var prev *TreeNode
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)

		if prev != nil {
			res = min(res, root.Val-prev.Val)
		}
		prev = root

		inorder(root.Right)
	}

	inorder(root)

	return res
}

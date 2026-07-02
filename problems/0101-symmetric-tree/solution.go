// https://leetcode.com/problems/symmetric-tree/description/
package symmetrictree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var check func(l, r *TreeNode) bool
	check = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}

		if l == nil || r == nil || l.Val != r.Val {
			return false
		}

		return check(l.Left, r.Right) && check(l.Right, r.Left)
	}

	return check(root.Left, root.Right)
}

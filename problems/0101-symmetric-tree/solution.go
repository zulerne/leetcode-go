// https://leetcode.com/problems/symmetric-tree/description/
package symmetrictree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetricIterative(root *TreeNode) bool {
	stack := []*TreeNode{root.Left, root.Right}

	for len(stack) >= 2 {
		r := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		l := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if r == l && r == nil {
			continue
		}
		if r == nil || l == nil || r.Val != l.Val {
			return false
		}
		stack = append(stack, r.Left, l.Right, r.Right, l.Left)
	}

	return len(stack) == 0
}

func isSymmetric(root *TreeNode) bool {
	var check func(l, r *TreeNode) bool
	check = func(l, r *TreeNode) bool {
		if r == l && r == nil {
			return true
		}
		if r == nil || l == nil || r.Val != l.Val {
			return false
		}
		return check(r.Right, l.Left) && check(r.Left, l.Right)
	}
	return check(root.Left, root.Right)
}

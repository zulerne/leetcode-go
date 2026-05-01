// https://leetcode.com/problems/diameter-of-binary-tree/description/
package diameterofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var diameter int
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := depth(root.Left)
		r := depth(root.Right)
		diameter = max(diameter, l+r)
		return max(l, r) + 1
	}
	depth(root)
	return diameter
}

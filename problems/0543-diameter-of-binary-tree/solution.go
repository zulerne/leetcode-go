// https://leetcode.com/problems/diameter-of-binary-tree/description/
package diameterofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var diameter int
	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := maxDepth(root.Left)
		r := maxDepth(root.Right)
		diameter = max(diameter, l+r)
		return max(l, r) + 1
	}
	maxDepth(root)

	return diameter
}

// https://leetcode.com/problems/binary-tree-level-order-traversal/description/
package binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		lvlSize := len(queue)
		lvl := make([]int, lvlSize)

		for i := range lvlSize {
			node := queue[0]
			queue = queue[1:]

			lvl[i] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, lvl)
	}

	return res
}

// https://leetcode.com/problems/binary-tree-level-order-traversal/description/
package binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	levels := make([][]int, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)

		for i := range levelSize {
			node := queue[0]
			queue = queue[1:]

			level[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levels = append(levels, level)
	}

	return levels
}

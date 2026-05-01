// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/
package binarytreezigzaglevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	levels := make([][]int, 0)
	queue := []*TreeNode{root}
	reverse := false

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)

		for i := range levelSize {
			node := queue[0]
			queue = queue[1:]

			if reverse {
				level[levelSize-1-i] = node.Val
			} else {
				level[i] = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levels = append(levels, level)
		reverse = !reverse
	}
	return levels
}

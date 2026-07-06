// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/
package binarytreezigzaglevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	var reverse bool

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		lvlSize := len(queue)
		lvl := make([]int, lvlSize)

		for i := range lvlSize {
			node := queue[0]
			queue = queue[1:]

			if reverse {
				lvl[lvlSize-1-i] = node.Val
			} else {
				lvl[i] = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, lvl)
		reverse = !reverse
	}

	return res
}

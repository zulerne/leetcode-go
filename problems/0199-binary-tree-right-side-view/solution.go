// https://leetcode.com/problems/binary-tree-right-side-view/description/
package binarytreerightsideview

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		lvlSize := len(queue)

		for i := range lvlSize {
			node := queue[0]
			queue = queue[1:]

			if i == lvlSize-1 {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return res
}

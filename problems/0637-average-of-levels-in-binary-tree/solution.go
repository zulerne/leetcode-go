// https://leetcode.com/problems/average-of-levels-in-binary-tree/description/
package averageoflevelsinbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		lvlSize := len(queue)
		var sum int

		for range lvlSize {
			node := queue[0]
			queue = queue[1:]

			sum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, float64(sum)/float64(lvlSize))
	}

	return res
}

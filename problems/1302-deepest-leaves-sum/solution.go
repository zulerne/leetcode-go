// https://leetcode.com/problems/deepest-leaves-sum/description/
package deepestleavessum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	var sum int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var levelSum int

		for range levelSize {
			node := queue[0]
			queue = queue[1:]

			levelSum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		sum = levelSum
	}

	return sum
}

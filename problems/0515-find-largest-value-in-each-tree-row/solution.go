// https://leetcode.com/problems/find-largest-value-in-each-tree-row/description/
package findlargestvalueineachtreerow

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		levelMax := math.MinInt
		for range levelSize {
			node := queue[0]
			queue = queue[1:]

			levelMax = max(levelMax, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, levelMax)
	}
	return result
}

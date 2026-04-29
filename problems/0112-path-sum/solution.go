// https://leetcode.com/problems/path-sum/description/
package pathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetDiff := targetSum - root.Val
	if targetDiff == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, targetDiff) || hasPathSum(root.Right, targetDiff)
}

type stackElem struct {
	node    *TreeNode
	current int
}

func hasPathSumIterative(root *TreeNode, targetSum int) bool {
	stack := []stackElem{{node: root, current: 0}}
	for len(stack) > 0 {
		el := stack[len(stack)-1]
		node, current := el.node, el.current
		stack = stack[:len(stack)-1]

		if node == nil {
			continue
		}

		current += node.Val

		if node.Left == nil && node.Right == nil && current == targetSum {
			return true
		}
		stack = append(stack, stackElem{node: node.Left, current: current})
		stack = append(stack, stackElem{node: node.Right, current: current})
	}
	return false
}

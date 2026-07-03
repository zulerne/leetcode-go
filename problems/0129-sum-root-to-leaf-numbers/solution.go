// https://leetcode.com/problems/sum-root-to-leaf-numbers/description/
package sumroottoleafnumbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var sum func(root *TreeNode, curSum int) int
	sum = func(root *TreeNode, curSum int) int {
		if root == nil {
			return 0
		}

		curSum = curSum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return curSum
		}

		return sum(root.Left, curSum) + sum(root.Right, curSum)
	}

	return sum(root, 0)
}

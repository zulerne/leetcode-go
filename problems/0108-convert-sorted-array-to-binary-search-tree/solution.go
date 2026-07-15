// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/
package convertsortedarraytobinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var build func(nums []int) *TreeNode
	build = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}

		mid := len(nums) / 2
		node := &TreeNode{Val: nums[mid]}
		node.Left = build(nums[:mid])
		node.Right = build(nums[mid+1:])

		return node
	}

	return build(nums)
}

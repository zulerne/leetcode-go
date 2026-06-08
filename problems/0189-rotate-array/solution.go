// https://leetcode.com/problems/rotate-array/description/
package rotatearray

import (
	"slices"
)

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n

	slices.Reverse(nums[:n-k])
	slices.Reverse(nums[n-k:])
	slices.Reverse(nums)
}

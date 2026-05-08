// Package maxofsumpairs
// https://leetcode.com/problems/max-number-of-k-sum-pairs
package maxofsumpairs

// Input: nums = [1,2,3,4], k = 5
// Output: 2
// Explanation: Starting with nums = [1,2,3,4]:
// - Remove numbers 1 and 4, then nums = [2,3]
// - Remove numbers 2 and 3, then nums = []
// TODO: Update
func maxOperations(nums []int, k int) int {
	op := 0

	//first, second := 0, 0

	for first := 0; first < len(nums); {
		firstNum := nums[first]
		for second := first + 1; second < len(nums); {
			secondNum := nums[second]

			if firstNum+secondNum == k {
				op++

			}

		}
	}

	return op
}

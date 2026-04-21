// https://leetcode.com/problems/minimum-size-subarray-sum/description/
package minimumsizesubarraysum

import (
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	minLength := math.MaxInt

	var sum int
	var beg int
	for end := range nums {
		sum += nums[end]

		for sum >= target {
			length := end - beg + 1
			minLength = min(minLength, length)
			sum -= nums[beg]
			beg++
		}
	}

	if minLength == math.MaxInt {
		return 0
	}
	return minLength
}

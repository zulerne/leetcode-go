// https://leetcode.com/problems/maximum-subarray/description/
package maximumsubarray

import "math"

func maxSubArray(nums []int) int {
	res := math.MinInt
	cur := 0

	for i := range nums {
		cur = max(nums[i], cur+nums[i])
		res = max(res, cur)
	}

	return res
}

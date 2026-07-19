// https://leetcode.com/problems/maximum-sum-circular-subarray/description/
package maximumsumcircularsubarray

import (
	"math"
)

func maxSubarraySumCircular(nums []int) int {
	var total, curMinSub, curMaxSub int
	minSub, maxSub := math.MaxInt, math.MinInt

	for _, n := range nums {
		curMinSub = min(n, curMinSub+n)
		minSub = min(minSub, curMinSub)

		curMaxSub = max(n, curMaxSub+n)
		maxSub = max(maxSub, curMaxSub)

		total += n
	}

	if minSub == total {
		return maxSub
	}

	return max(maxSub, total-minSub)
}

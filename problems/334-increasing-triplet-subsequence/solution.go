// Package tripletsubseq
// https://leetcode.com/problems/increasing-triplet-subsequence
package tripletsubseq

import (
	"math"
)

func increasingTriplet(nums []int) bool {
	first := math.MaxInt
	second := math.MaxInt

	for i := 0; i < len(nums); i++ {
		n := nums[i]

		if n <= first {
			first = n
		} else if n <= second {
			second = n
		} else {
			return true
		}
	}

	return false
}

// https://leetcode.com/problems/permutations/description/
package permutations

import "slices"

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	seen := make([]bool, len(nums))
	acc := make([]int, len(nums))

	var backtrack func(pos int)
	backtrack = func(pos int) {
		if pos == len(nums) {
			res = append(res, slices.Clone(acc))
			return
		}

		for i := range nums {
			if seen[i] {
				continue
			}
			acc[pos] = nums[i]
			seen[i] = true
			backtrack(pos + 1)
			seen[i] = false
		}
	}

	backtrack(0)

	return res
}

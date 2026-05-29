// https://leetcode.com/problems/permutations/description/
package permutations

import (
	"slices"
)

func permute(nums []int) [][]int {
	var res [][]int
	var backtrack func(acc []int)

	backtrack = func(acc []int) {
		if len(acc) == len(nums) {
			tmp := make([]int, len(acc))
			copy(tmp, acc)
			res = append(res, tmp)
			return
		}

		for _, num := range nums {
			if !slices.Contains(acc, num) {
				acc = append(acc, num)
				backtrack(acc)
				acc = acc[:len(acc)-1]
			}
		}
	}

	backtrack([]int{})

	return res
}

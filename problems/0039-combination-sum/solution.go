// https://leetcode.com/problems/combination-sum/description/
package combinationsum

import "slices"

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	acc := []int{}

	var backtrack func(pos, diff int)
	backtrack = func(pos, diff int) {
		if diff == 0 {
			res = append(res, slices.Clone(acc))
			return
		}

		for i := pos; i < len(candidates); i++ {
			c := candidates[i]
			if diff >= c {
				acc = append(acc, c)
				backtrack(i, diff-c)
				acc = acc[:len(acc)-1]
			}
		}
	}

	backtrack(0, target)

	return res
}

// https://leetcode.com/problems/subsets/description/
package subsets

func subsets(nums []int) [][]int {
	var res [][]int
	var backtrack func(acc []int, cur int)

	backtrack = func(acc []int, cur int) {
		tmp := make([]int, len(acc))
		copy(tmp, acc)
		res = append(res, tmp)

		for i := cur; i < len(nums); i++ {
			acc = append(acc, nums[i])
			backtrack(acc, i+1)
			acc = acc[:len(acc)-1]
		}
	}

	backtrack([]int{}, 0)

	return res
}

// https://leetcode.com/problems/combination-sum-iii/description/
package combinationsumiii

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var backtrack func(acc []int, sum, cur int)

	backtrack = func(acc []int, sum, cur int) {
		if len(acc) == k && sum == n {
			accCopy := make([]int, len(acc))
			copy(accCopy, acc)
			res = append(res, accCopy)
			return
		}

		for i := cur; i < 10; i++ {
			acc = append(acc, i)
			backtrack(acc, sum+i, i+1)
			acc = acc[:len(acc)-1]
		}
	}

	backtrack([]int{}, 0, 1)
	return res
}

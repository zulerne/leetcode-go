// https://leetcode.com/problems/combinations/description/
package combinations

func combine(n int, k int) [][]int {
	var result [][]int
	var backtrack func(acc []int, i int)

	backtrack = func(acc []int, cur int) {
		if len(acc) == k {
			tmp := make([]int, len(acc))
			copy(tmp, acc)
			result = append(result, tmp)
			return
		}

		for i := cur; i < n+1; i++ {
			acc = append(acc, i)
			backtrack(acc, i+1)
			acc = acc[:len(acc)-1]
		}
	}

	backtrack([]int{}, 1)

	return result
}

// https://leetcode.com/problems/combinations/description/
package combinations

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	acc := make([]int, k)

	var backtrack func(start, pos int)
	backtrack = func(start, pos int) {
		if pos == k {
			r := make([]int, k)
			copy(r, acc)
			res = append(res, r)
			return
		}

		need := k - pos
		for cur := start + 1; cur <= n-need+1; cur++ {
			if pos+1 == k+1 {
				break
			}
			acc[pos] = cur
			backtrack(cur, pos+1)
		}
	}

	backtrack(0, 0)

	return res
}

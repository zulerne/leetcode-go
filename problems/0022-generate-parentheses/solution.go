// https://leetcode.com/problems/generate-parentheses/description/
package generateparentheses

func generateParenthesis(n int) []string {
	var res []string
	acc := make([]byte, n*2)

	var backtrack func(pos, open, closed int)
	backtrack = func(pos, open, closed int) {
		if pos == n*2 {
			res = append(res, string(acc))
			return
		}
		if open < n {
			acc[pos] = '('
			backtrack(pos+1, open+1, closed)
		}
		if closed < open {
			acc[pos] = ')'
			backtrack(pos+1, open, closed+1)
		}
	}

	backtrack(0, 0, 0)

	return res
}

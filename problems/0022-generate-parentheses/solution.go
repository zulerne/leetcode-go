// https://leetcode.com/problems/generate-parentheses/description/
package generateparentheses

func generateParenthesis(n int) []string {
	var res []string
	var backtrack func(acc []byte, open, closed int)

	backtrack = func(acc []byte, open, closed int) {
		if len(acc) == n*2 {
			tmp := make([]byte, len(acc))
			copy(tmp, acc)
			res = append(res, string(tmp))
		}

		if open < n {
			acc = append(acc, '(')
			backtrack(acc, open+1, closed)
			acc = acc[:len(acc)-1]
		}
		if closed < open {
			acc = append(acc, ')')
			backtrack(acc, open, closed+1)
			acc = acc[:len(acc)-1]
		}
	}

	backtrack([]byte{}, 0, 0)

	return res
}

// https://leetcode.com/problems/valid-parentheses/description/
package validparentheses

func isValid(s string) bool {
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	var stack []byte

	for i := range s {
		switch b := s[i]; b {
		case '(', '[', '{':
			stack = append(stack, b)
		case ')', ']', '}':
			lastIdx := len(stack) - 1
			if lastIdx < 0 {
				return false
			}
			if pairs[b] != stack[lastIdx] {
				return false
			}
			stack = stack[:lastIdx]
		}
	}

	return len(stack) == 0
}

// https://leetcode.com/problems/valid-parentheses/description/
package validparentheses

func isValid(s string) bool {
	m := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0, len(s))

	for i := range s {
		switch b := s[i]; b {
		case '(', '[', '{':
			stack = append(stack, b)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if m[b] != last {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

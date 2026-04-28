// https://leetcode.com/problems/removing-stars-from-a-string/description/
package removingstarsfromastring

func removeStars(s string) string {
	var stack []byte

	for i := range s {
		b := s[i]
		if b == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, b)
		}
	}

	return string(stack)
}

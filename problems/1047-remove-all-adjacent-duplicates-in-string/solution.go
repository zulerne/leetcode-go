// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/description/
package removealladjacentduplicatesinstring

func removeDuplicates(s string) string {
	var stack []byte

	for i := range s {
		b := s[i]
		length := len(stack)
		if length > 0 && b == stack[length-1] {
			stack = stack[:length-1]
			continue
		}
		stack = append(stack, b)
	}

	return string(stack)
}

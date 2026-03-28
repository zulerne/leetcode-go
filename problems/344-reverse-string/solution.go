// Package template - template for new solutions
// https://leetcode.com/problems/reverse-string/description/
package reversestring

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

// Package template - template for new solutions
// https://leetcode.com/problems/valid-palindrome/description/
package validpalindrome

func isAlphanumeric(b byte) bool {
	return (b >= 48 && b <= 57) || (b >= 65 && b <= 90) || (b >= 97 && b <= 122)
}

func toLower(b byte) byte {
	if b >= 65 && b <= 90 {
		return b + 32
	}
	return b
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		switch {
		case !isAlphanumeric(s[i]):
			i++
		case !isAlphanumeric(s[j]):
			j--
		case toLower(s[i]) != toLower(s[j]):
			return false
		default:
			i, j = i+1, j-1
		}
	}
	return true
}

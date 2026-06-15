// https://leetcode.com/problems/valid-palindrome/description/
package validpalindrome

func isAlphanumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		for i < j && !isAlphanumeric(s[i]) {
			i++
		}
		for i < j && !isAlphanumeric(s[j]) {
			j--
		}

		if i < j {
			if toLower(s[i]) != toLower(s[j]) {
				return false
			}
			i++
			j--
		}
	}

	return true
}

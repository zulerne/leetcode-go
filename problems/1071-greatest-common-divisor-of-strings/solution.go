// Package stringsdivisor
// https://leetcode.com/problems/greatest-common-divisor-of-strings/description
package stringsdivisor

import (
	"strings"
)

// O(min(N, M) *(N+M))
func gcdOfStringsFirst(str1 string, str2 string) string {
	var gcd string
	n, m := len(str1), len(str2)
	minLen := min(len(str1), len(str2))

	for i := 1; i < minLen+1; i++ {
		prefix := str1[:i]

		if strings.Repeat(prefix, n/i) == str1 &&
			strings.Repeat(prefix, m/i) == str2 {
			gcd = prefix
		}
	}

	return gcd
}

// gcdOfStrings has O(N + M) complexity
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	length := gcd(len(str1), len(str2))

	return str1[:length]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

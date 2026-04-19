// Package reversevowels
// https://leetcode.com/problems/3/description
package reversevowels

import (
	"unicode"
)

func reverseVowels(s string) string {
	runes := []rune(s)
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isVowel(runes[l]) {
			l++
		}

		for r > l && !isVowel(runes[r]) {
			r--
		}

		runes[l], runes[r] = runes[r], runes[l]
		l++
		r--
	}

	return string(runes)
}

func isVowel(b rune) bool {
	switch unicode.ToLower(b) {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

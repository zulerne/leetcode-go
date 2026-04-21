// Package reversewords
// https://leetcode.com/problems/reverse-words-in-a-string
package reversewords

import (
	"slices"
	"strings"
	"unicode"
)

func reverseWordsManual(s string) string {
	s = strings.TrimSpace(s)
	words := make([]string, 0)
	b := strings.Builder{}

	for i, r := range s {
		if !unicode.IsSpace(r) {
			b.WriteRune(r)
		}
		if unicode.IsSpace(r) || i == len(s)-1 {
			if b.Len() > 0 {
				words = append(words, b.String())
				b.Reset()
			}
		}
	}

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func reverseWords(s string) string {
	words := strings.Fields(s)
	slices.Reverse(words)
	return strings.Join(words, " ")
}

// Package mergestrings
// https://leetcode.com/problems/merge-strings-alternately/description
package mergestrings

import (
	"strings"
	"unsafe"
)

func mergeAlternatelyFirst(word1 string, word2 string) string {
	res := make([]byte, 0, len(word1)+len(word2))
	minLen := min(len(word1), len(word2))

	for i := 0; i < minLen; i++ {
		res = append(res, word1[i], word2[i])
	}
	res = append(res, word1[minLen:]...)
	res = append(res, word2[minLen:]...)

	return unsafe.String(unsafe.SliceData(res), len(res))
}

func mergeAlternately(word1 string, word2 string) string {
	n, m := len(word1), len(word2)
	minLen := min(n, m)

	b := strings.Builder{}
	b.Grow(n + m)

	for i := 0; i < minLen; i++ {
		b.WriteByte(word1[i])
		b.WriteByte(word2[i])
	}
	if n > m {
		b.WriteString(word1[minLen:])
	} else {
		b.WriteString(word2[minLen:])
	}

	return b.String()
}

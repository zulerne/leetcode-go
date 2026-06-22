// https://leetcode.com/problems/word-pattern/description/
package wordpattern

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	pMap := [26]string{}
	sMap := make(map[string]byte)
	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	for i := range pattern {
		pByte := pattern[i] - 'a'
		sWord := words[i]

		if pMap[pByte] != "" && pMap[pByte] != sWord {
			return false
		}
		if v, ok := sMap[sWord]; ok && v != pByte {
			return false
		}

		pMap[pByte] = sWord
		sMap[sWord] = pByte
	}

	return true
}

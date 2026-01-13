// Package issubseq
// https://leetcode.com/problems/is-subsequence
package issubseq

func isSubsequenceFirst(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	for i, j := 0, 0; i < len(s) && j < len(t); {
		cur := s[i]

		for j < len(t) {
			if t[j] == cur {
				i++
				j++
				break
			}
			j++
		}

		if i == len(s) {
			return true
		}
	}

	return false
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	for i, j := 0, 0; i < len(t); i++ {
		if t[i] == s[j] {
			j++
		}
		if j == len(s) {
			return true
		}
	}

	return false
}

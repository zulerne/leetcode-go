// https://leetcode.com/problems/reverse-words-in-a-string/description/
package reversewordsinastring

import (
	"strings"
)

func reverseWords(s string) string {
	var res strings.Builder
	end := len(s) - 1

	res.Grow(len(s))

	for end >= 0 {
		for end >= 0 && s[end] == ' ' {
			end--
		}
		if end < 0 {
			break
		}

		start := end
		for start >= 0 && s[start] != ' ' {
			start--
		}

		if res.Len() > 0 {
			res.WriteByte(' ')
		}
		res.WriteString(s[start+1 : end+1])

		end = start
	}
	return res.String()
}

// https://leetcode.com/problems/length-of-last-word/description/
package lengthoflastword

func lengthOfLastWord(s string) int {
	var res int

	i := len(s) - 1
	for ; i >= 0 && s[i] == ' '; i-- {
	}
	for ; i >= 0 && s[i] != ' '; i-- {
		res++
	}

	return res
}

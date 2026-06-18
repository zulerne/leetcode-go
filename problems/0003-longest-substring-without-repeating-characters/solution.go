// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	var res int
	chars := [256]bool{}
	begin := 0

	for end := 0; end < len(s); end++ {
		for chars[s[end]] {
			chars[s[begin]] = false
			begin++
		}

		chars[s[end]] = true

		res = max(res, end-begin+1)
	}
	return res
}

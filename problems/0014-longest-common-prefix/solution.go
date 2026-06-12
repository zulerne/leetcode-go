// https://leetcode.com/problems/longest-common-prefix/description/
package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	first := strs[0]

	for i := range len(first) {
		cur := first[i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != cur {
				return first[:i]
			}
		}
	}

	return first
}

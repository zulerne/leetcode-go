// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
package findtheindexofthefirstoccurrenceinastring

func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)

	for i := 0; i <= n-m; i++ {
		if haystack[i:i+m] == needle {
			return i
		}
	}

	return -1
}

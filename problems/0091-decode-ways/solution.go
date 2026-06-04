// https://leetcode.com/problems/decode-ways/description/
package decodeways

func numDecodings(s string) int {
	cache := make(map[string]int)
	var inner func(s string) int
	inner = func(s string) int {
		if c, ok := cache[s]; ok {
			return c
		}

		if len(s) == 0 {
			return 1
		}

		if s[0] == '0' {
			return 0
		}

		if len(s) == 1 {
			return 1
		}

		if s[:2] > "26" {
			cache[s] = inner(s[1:])
			return cache[s]
		}

		cache[s] = inner(s[1:]) + inner(s[2:])
		return cache[s]
	}

	return inner(s)
}

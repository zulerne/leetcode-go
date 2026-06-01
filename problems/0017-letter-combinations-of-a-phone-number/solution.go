// https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
package lettercombinationsofaphonenumber

func letterCombinations(digits string) []string {
	var res []string

	letters := make(map[byte][]byte)
	letters['2'] = []byte{'a', 'b', 'c'}
	letters['3'] = []byte{'d', 'e', 'f'}
	letters['4'] = []byte{'g', 'h', 'i'}
	letters['5'] = []byte{'j', 'k', 'l'}
	letters['6'] = []byte{'m', 'n', 'o'}
	letters['7'] = []byte{'p', 'q', 'r', 's'}
	letters['8'] = []byte{'t', 'u', 'v'}
	letters['9'] = []byte{'w', 'x', 'y', 'z'}

	var backtrack func(acc []byte, cur int)
	backtrack = func(acc []byte, cur int) {
		if cur == len(digits) {
			res = append(res, string(acc))
			return
		}
		for _, v := range letters[digits[cur]] {
			acc = append(acc, v)
			backtrack(acc, cur+1)
			acc = acc[:len(acc)-1]
		}
	}

	backtrack([]byte{}, 0)

	return res
}

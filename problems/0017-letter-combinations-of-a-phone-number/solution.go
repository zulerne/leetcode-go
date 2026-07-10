// https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
package lettercombinationsofaphonenumber

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	var letters = [10]string{
		"",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}

	res := make([]string, 0)
	path := make([]byte, len(digits))

	var backtrack func(pos int)
	backtrack = func(pos int) {
		if pos == len(digits) {
			res = append(res, string(path))
			return
		}

		digitLetters := letters[digits[pos]-'0']
		for i := range digitLetters {
			path[pos] = digitLetters[i]
			backtrack(pos + 1)
		}
	}

	backtrack(0)

	return res
}

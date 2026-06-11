// https://leetcode.com/problems/roman-to-integer/description/
package romantointeger

func romanValue(b byte) int {
	switch b {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

func romanToInt(s string) int {
	var res int

	for i := range s {
		cur := romanValue(s[i])
		if i+1 < len(s) && cur < romanValue(s[i+1]) {
			res -= cur
		} else {
			res += cur
		}
	}

	return res
}

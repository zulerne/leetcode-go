// https://leetcode.com/problems/backspace-string-compare/description/
package backspacestringcompare

func backspaceCompare(s string, t string) bool {
	sc, tc := 0, 0

	for si, ti := len(s)-1, len(t)-1; ; si, ti = si-1, ti-1 {
		for si >= 0 && (s[si] == '#' || sc > 0) {
			if s[si] == '#' {
				sc++
			} else {
				sc--
			}
			si--
		}

		for ti >= 0 && (t[ti] == '#' || tc > 0) {
			if t[ti] == '#' {
				tc++
			} else {
				tc--
			}
			ti--
		}

		if si < 0 && ti < 0 {
			return true
		}
		if si < 0 || ti < 0 {
			return false
		}
		if s[si] != t[ti] {
			return false
		}
	}
}

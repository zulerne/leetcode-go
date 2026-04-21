// https://leetcode.com/problems/isomorphic-strings/description/
package isomorphicstrings

func isIsomorphic(s string, t string) bool {
	sMap, tMap := make(map[byte]byte), make(map[byte]byte)
	for i := range len(s) {
		sc, tc := s[i], t[i]
		if v, ok := sMap[sc]; ok && v != tc {
			return false
		}
		if v, ok := tMap[tc]; ok && v != sc {
			return false
		}
		sMap[sc], tMap[tc] = tc, sc
	}
	return true
}

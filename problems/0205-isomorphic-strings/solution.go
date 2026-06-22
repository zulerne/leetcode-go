// https://leetcode.com/problems/isomorphic-strings/description/
package isomorphicstrings

func isIsomorphic(s string, t string) bool {
	n := len(s)
	smap := [256]byte{}
	tmap := [256]byte{}

	for i := range n {
		sChar, tChar := s[i], t[i]

		if smap[sChar] != 0 && smap[sChar] != tChar {
			return false
		}
		if tmap[tChar] != 0 && tmap[tChar] != sChar {
			return false
		}

		smap[sChar], tmap[tChar] = tChar, sChar
	}

	return true
}

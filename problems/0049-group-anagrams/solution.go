// https://leetcode.com/problems/group-anagrams/description/
package groupanagrams

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0, len(strs))
	m := make(map[string][]string)

	for _, s := range strs {
		letters := []byte(s)
		slices.Sort(letters)

		m[string(letters)] = append(m[string(letters)], s)
	}

	for _, v := range m {
		res = append(res, v)
	}

	return res
}

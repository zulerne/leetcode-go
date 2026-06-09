// https://leetcode.com/problems/h-index/description/
package hindex

import (
	"slices"
)

func hIndex(citations []int) int {
	var res int
	n := len(citations)

	slices.Sort(citations)

	for i, v := range citations {
		if v >= n-i {
			res = n - i
			break
		}
	}

	return res
}

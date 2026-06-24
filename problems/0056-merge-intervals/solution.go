// https://leetcode.com/problems/merge-intervals/description/
package mergeintervals

import (
	"slices"
)

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	res := make([][]int, 0, n)

	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	for i := 0; i < n; i++ {
		a, b := intervals[i][0], intervals[i][1]

		for i+1 < n && b >= intervals[i+1][0] {
			b = max(b, intervals[i+1][1])
			i++
		}
		res = append(res, []int{a, b})
	}

	return res
}

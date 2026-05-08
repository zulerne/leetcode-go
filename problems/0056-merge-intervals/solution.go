// https://leetcode.com/problems/merge-intervals/description/
package mergeintervals

import "sort"

func mergeAlternative(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)

	for i := 1; i <= len(intervals); i++ {
		start := intervals[i-1][0]
		end := intervals[i-1][1]
		for i < len(intervals) && intervals[i][0] <= end {
			end = max(end, intervals[i][1])
			i++
		}
		result = append(result, []int{start, end})
	}

	return result
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		if intervals[i][0] <= last[1] {
			result[len(result)-1][1] = max(result[len(result)-1][1], intervals[i][1])
		} else {
			result = append(result, intervals[i])
		}
	}

	return result
}

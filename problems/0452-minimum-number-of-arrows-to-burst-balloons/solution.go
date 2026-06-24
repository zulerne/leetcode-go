// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/description/
package minimumnumberofarrowstoburstballoons

import (
	"slices"
)

func findMinArrowShots(points [][]int) int {
	slices.SortFunc(points, func(a, b []int) int {
		return a[1] - b[1]
	})

	arrows := 1
	firstEnd := points[0][1]

	for i := 1; i < len(points); i++ {
		if points[i][0] > firstEnd {
			arrows++
			firstEnd = points[i][1]
		}
	}
	return arrows
}

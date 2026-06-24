// https://leetcode.com/problems/insert-interval/description/
package insertinterval

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	i := 0

	res := make([][]int, 0)

	for ; i < n && intervals[i][1] < newInterval[0]; i++ {
		res = append(res, intervals[i])
	}

	for ; i < n && intervals[i][0] <= newInterval[1]; i++ {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
	}
	res = append(res, newInterval)

	for ; i < n; i++ {
		res = append(res, intervals[i])
	}

	return res
}

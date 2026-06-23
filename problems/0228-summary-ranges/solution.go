// https://leetcode.com/problems/summary-ranges/description/
package summaryranges

import "fmt"

func summaryRanges(nums []int) []string {
	n := len(nums)
	res := make([]string, 0)

	for i := 0; i < n; {
		start := nums[i]

		for i+1 < n && nums[i+1] == nums[i]+1 {
			i++
		}

		if start == nums[i] {
			res = append(res, fmt.Sprintf("%v", start))
		} else {
			res = append(res, fmt.Sprintf("%v->%v", start, nums[i]))
		}

		i++
	}

	return res
}

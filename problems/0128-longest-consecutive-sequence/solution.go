// https://leetcode.com/problems/longest-consecutive-sequence/description/
package longestconsecutivesequence

func longestConsecutive(nums []int) int {
	m := make(map[int]int, len(nums))
	res := 0

	for _, n := range nums {
		if _, ok := m[n]; ok {
			continue
		}

		left := m[n-1]
		right := m[n+1]

		cur := left + right + 1
		m[n] = cur
		res = max(res, cur)

		m[n-left] = cur
		m[n+right] = cur
	}

	return res
}

// func longestConsecutive(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	slices.Sort(nums)

// 	res := 0
// 	cur := 1
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i]-1 == nums[i-1] {
// 			cur++
// 		} else if nums[i] == nums[i-1] {
// 			continue
// 		} else {
// 			res = max(res, cur)
// 			cur = 1
// 		}
// 	}

// 	res = max(res, cur)

// 	return res
// }

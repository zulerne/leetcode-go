// https://leetcode.com/problems/3sum/description/
package threesum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	length := len(nums)

	sort.Ints(nums)
	result := make([][]int, 0)

	for i, v := range nums {
		target := -v

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for l, r := i+1, length-1; l < r; {
			sum := nums[l] + nums[r]
			if sum < target {
				l++
				continue
			} else if sum > target {
				r--
				continue
			} else {
				s := []int{v, nums[l], nums[r]}
				result = append(result, s)
				l, r = l+1, r-1

				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for r > l && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}

	return result
}

// https://leetcode.com/problems/3sum/description/
package threesum

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	n := len(nums)

	slices.Sort(nums)

	for i, v := range nums {
		target := -v

		if i > 0 && nums[i-1] == v {
			continue
		}

		for j, k := i+1, n-1; j < k; {
			sum := nums[j] + nums[k]
			if sum < target {
				j++
			} else if sum > target {
				k--
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--

				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
			}
		}
	}

	return res
}

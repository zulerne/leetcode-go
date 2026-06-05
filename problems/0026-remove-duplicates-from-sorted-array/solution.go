// Package template - template for new solutions
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	j := 0
	for i := range nums {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}

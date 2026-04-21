// https://leetcode.com/problems/binary-search/description/
package binarysearch

func search(nums []int, target int) int {
	low := 0
	high := len(nums)

	for mid := low + (high-low)/2; low < high; mid = low + (high-low)/2 {
		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid
		} else {
			return mid
		}
	}
	return -1
}

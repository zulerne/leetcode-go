// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
package findfirstandlastpositionofelementinsortedarray

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}

	left, right := 0, len(nums)-1
	ans := -1
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			ans = mid
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if ans == -1 {
		return res
	}
	res[0] = ans

	left, right = 0, len(nums)-1
	ans = -1
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			ans = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	res[1] = ans

	return res
}

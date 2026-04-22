// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/description/
package longestsubarrayof1safterdeletingoneelement

func longestSubarray(nums []int) int {
	res := 0
	zeros := 0
	begin := 0
	for end := range nums {
		if nums[end] == 0 {
			zeros++
		}

		for zeros > 1 {
			if nums[begin] == 0 {
				zeros--
			}
			begin++
		}
		res = max(res, end-begin)
	}
	return res
}

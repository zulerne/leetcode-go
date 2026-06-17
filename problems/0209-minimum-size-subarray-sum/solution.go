// https://leetcode.com/problems/minimum-size-subarray-sum/description/
package minimumsizesubarraysum

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	res := n + 1
	var begin, sum int

	for end := range n {
		sum += nums[end]
		for sum >= target {
			res = min(res, end-begin+1)
			sum -= nums[begin]
			begin++
		}
	}

	if res == n+1 {
		return 0
	}
	return res
}

// https://leetcode.com/problems/maximum-average-subarray-i/description/
package maximumaveragesubarrayi

func findMaxAverage(nums []int, k int) float64 {
	var sum int
	for _, v := range nums[:k] {
		sum += v
	}

	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}

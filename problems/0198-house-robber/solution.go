// https://leetcode.com/problems/house-robber/description/
package houserobber

func rob(nums []int) int {
	a, b := 0, 0
	for _, n := range nums {
		tmp := max(a+n, b)
		a = b
		b = tmp
	}
	return b
}

// func rob(nums []int) int {
// 	n := len(nums)
// 	dp := make([]int, n+1)
// 	dp[1] = nums[0]

// 	for i := 2; i < len(dp); i++ {
// 		dp[i] = max(nums[i-1]+dp[i-2], dp[i-1])
// 	}

// 	return dp[n]
// }

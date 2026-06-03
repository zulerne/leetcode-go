// https://leetcode.com/problems/min-cost-climbing-stairs/description/
package mincostclimbingstairs

import "fmt"

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)

	for i := 2; i < len(dp); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	fmt.Printf("dp: %v\n", dp)
	return dp[n]
}

// func minCostClimbingStairs(cost []int) int {
// 	n := len(cost)
// 	cache := make([]int, n+1)
// 	for i := range cache {
// 		cache[i] = -1
// 	}

// 	var inner func(n int) int
// 	inner = func(n int) int {
// 		if n <= 1 {
// 			return 0
// 		}
// 		if cache[n] != -1 {
// 			return cache[n]
// 		}

// 		cache[n] = min(inner(n-1)+cost[n-1], inner(n-2)+cost[n-2])

// 		return cache[n]
// 	}

// 	return inner(n)
// }

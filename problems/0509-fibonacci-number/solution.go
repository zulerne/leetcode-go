// https://leetcode.com/problems/fibonacci-number/description/
package fibonaccinumber

func fib(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for range n - 1 {
		a, b = b, a+b
	}
	return b
}

// func fib(n int) int {
// 	if n <= 1 {
// 		return n
// 	}

// 	dp := make([]int, n+1)
// 	dp[0] = 0
// 	dp[1] = 1

// 	for i := 2; i < len(dp); i++ {
// 		dp[i] = dp[i-1] + dp[i-2]
// 	}

// 	return dp[n]
// }

// func fib(n int) int {
// 	cache := make([]int, n+1)

// 	var inner func(n int) int
// 	inner = func(n int) int {
// 		if n <= 1 {
// 			return n
// 		}
// 		if cache[n] != 0 {
// 			return cache[n]
// 		}
// 		cache[n] = inner(n-1) + inner(n-2)
// 		return cache[n]
// 	}

// 	return inner(n)
// }

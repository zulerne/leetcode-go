// https://leetcode.com/problems/unique-paths/description/
package uniquepaths

func uniquePaths(m int, n int) int {
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, m)
	}

	for i := range n {
		for j := range m {
			dp[i][j] = 1
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[n-1][m-1]

}

// https://leetcode.com/problems/minimum-path-sum/description/
package minimumpathsum

func minPathSum(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + grid[i-1][0]
	}
	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j-1]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = min(dp[i][j-1]+grid[i][j-1], dp[i-1][j]+grid[i-1][j])
		}
	}

	return dp[n-1][m-1] + grid[n-1][m-1]
}

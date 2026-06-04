// https://leetcode.com/problems/edit-distance/description/
package editdistance

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}
	for j := 1; j < m+1; j++ {
		dp[0][j] = dp[0][j-1] + 1
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(
					dp[i-1][j-1],
					dp[i-1][j],
					dp[i][j-1],
				)
			}
		}
	}
	return dp[n][m]
}

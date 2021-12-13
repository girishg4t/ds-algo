package move

func Move(m, n int) uint8 {
	dp := make([][]uint8, m)
	for i := range dp {
		dp[i] = make([]uint8, n)
	}

	for row := 0; row < m; row++ {
		dp[row][0] = 1
	}

	for col := 0; col < n; col++ {
		dp[0][col] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = 1
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

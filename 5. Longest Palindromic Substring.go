package main

//DP
func longestPalindrome(s string) string {
	ans := ""
	n := len(s)
	var dp [][]bool
	for i := 0; i < n; i++ {
		var row []bool
		for j := 0; j < n; j++ {
			row = append(row, false)
		}
		dp = append(dp, row)
	}
	for l := 0; l < n; l++ {
		for i := 0; i < n; i++ {
			j := i + l
			if j >= n {
				break
			}
			if l == 0 {
				dp[i][j] = true
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = true
				}
			} else {
				if dp[i+1][j-1] && s[i] == s[j] {
					dp[i][j] = true
				}
			}
			if dp[i][j] && l+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}

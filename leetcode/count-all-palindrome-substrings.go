package leetcode

func CountAllPalindromeSubstrings(s string) int {

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	palindromic := 0

	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if i-j > 1 {
				dp[j][i] = dp[j+1][i-1] && s[j] == s[i]
			} else {
				dp[j][i] = s[j] == s[i]
			}

			if dp[j][i] {
				palindromic++
			}
		}
	}

	return palindromic
}

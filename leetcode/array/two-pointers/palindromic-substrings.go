package two_pointers

// https://leetcode.com/problems/palindromic-substrings/

func CountSubstrings(s string) int {

	count := 0

	for i := 0; i < len(s); i++ {

		// odd-length palindromes, single character center
		count += extendPalindrome(s, i, i)

		// even-length palindromes, consecutive characters center
		count += extendPalindrome(s, i, i+1)
	}

	return count
}

func extendPalindrome(s string, l, r int) int {

	res := 0

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
		res++
	}

	return res
}

func CountSubstringsDP(s string) int {

	if len(s) == 0 {
		return 0
	}

	ans, n := 0, len(s)

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = true
		ans++
	}

	for i := 0; i < n-1; i++ {
		dp[i][i+1] = s[i] == s[i+1]
		if dp[i][i+1] {
			ans++
		}
	}

	for lenS := 3; lenS <= n; lenS++ {
		j := lenS - 1
		for i := 0; j < n; i++ {
			dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			if dp[i][j] {
				ans++
			}
			j++
		}
	}

	return ans
}

func CountSubstringsDP2(s string) int {

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	ans := 0

	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if i-j > 1 {
				dp[j][i] = dp[j+1][i-1] && s[i] == s[j]
			} else {
				dp[j][i] = s[i] == s[j]
			}

			if dp[j][i] {
				ans++
			}
		}
	}

	return ans
}

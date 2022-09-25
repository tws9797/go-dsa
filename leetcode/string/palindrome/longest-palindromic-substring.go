package palindrome

// https://leetcode.com/problems/longest-palindromic-substring/

func LongestPalindrome(s string) string {

	if s == "" {
		return ""
	}

	var max string

	for i := 0; i < len(s); i++ {
		max = getLongestPalindrome(s, i, i, max)
		max = getLongestPalindrome(s, i, i+1, max)
	}

	return max
}

func getLongestPalindrome(s string, l, r int, max string) string {

	var sub string

	for l >= 0 && r < len(s) && s[r] == s[l] {
		sub = s[l : r+1]
		l--
		r++
	}

	if len(sub) > len(max) {
		max = sub
	}

	return max
}

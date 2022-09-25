package essentials

// https://leetcode.com/problems/longest-palindrome/

func LongestPalindrome(s string) int {

	m := map[byte]int{}

	if s == "" {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	count := 0

	for i := 0; i < len(s); i++ {
		m[s[i]]++
		if m[s[i]]%2 == 0 {
			count = count + 2
		}
	}

	for _, v := range m {
		if v%2 == 1 {
			count++
			break
		}
	}

	return count
}

func longestPalindromeGreedy(s string) int {
	count := make([]int, 128)

	for _, r := range s {
		count[r]++
	}

	ans := 0

	for _, v := range count {
		ans += v / 2 * 2
		if ans%2 == 0 && v%2 == 1 {
			ans++
		}
	}

	return ans
}
package sliding_window

import "math"

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func LengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	l, r := 0, 0
	maxLen := math.MinInt
	m := map[byte]int{}

	for r < len(s) {

		m[s[r]]++

		// Check if the current character already exists in the hash set
		for m[s[r]] > 1 {
			m[s[l]]--
			l++
		}

		r++
		if r-l > maxLen {
			maxLen = r - l
		}
	}

	return maxLen
}

func LengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}

	ans := 0
	m := map[byte]int{}
	l := 0

	for r, b := range []byte(s) {

		if prevIdx, ok := m[b]; ok {

			// To prevent scenario like "abba" where a is being included when reach the last index
			// To prevent l to move backwards
			if prevIdx > l {
				l = prevIdx
			}
		}

		if r-l+1 > ans {
			ans = r - l + 1
		}

		m[b] = r + 1
	}

	return ans
}

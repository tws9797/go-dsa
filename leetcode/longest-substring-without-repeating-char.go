package leetcode

import "math"

func LengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	l, r := 0, 0
	maxLen := math.MinInt
	m := map[byte]int{}

	for r < len(s) {

		m[s[r]]++

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

	for i, r := range []byte(s) {

		if prevIdx, ok := m[r]; ok {

			// To prevent scenario like "abba" where a is being included when reach the last index
			if prevIdx > l {
				l = prevIdx
			}
		}

		if i-l+1 > ans {
			ans = i - l + 1
		}

		m[r] = i + 1
	}

	return ans
}

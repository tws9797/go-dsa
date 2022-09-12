package leetcode

func LengthOfLongestSubstringTwoDistinct(s string) int {

	// Sliding window approach
	m := map[byte]int{}

	// Set left and right pointer in the position 0
	left, right := 0, 0
	max := 2

	if len(s) < 3 {
		return len(s)
	}

	for right < len(s) {
		m[s[right]]++

		for len(m) > 2 {
			m[s[left]]--
			if m[s[left]] == 0 {
				delete(m, s[left])
			}
			left++
		}

		right++
		if right-left > max {
			max = right - left
		}
	}

	return max
}

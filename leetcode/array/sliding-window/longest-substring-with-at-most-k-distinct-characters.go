package sliding_window

// https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/

func LengthOfLongestSubstringKDistinct(s string, k int) int {

	if k > len(s) {
		return len(s)
	}

	m := map[byte]int{}
	l := 0
	longest := k

	for i := 0; i < len(s); i++ {

		m[s[i]]++

		for len(m) > k {

			m[s[l]]--
			if m[s[l]] == 0 {
				delete(m, s[l])
			}
			l++
		}

		if i-l+1 > longest {
			longest = i - l + 1
		}
	}

	return longest
}

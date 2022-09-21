package essentials

// https://leetcode.com/problems/ransom-note/

func canConstruct(ransomNote string, magazine string) bool {

	if len(ransomNote) > len(magazine) {
		return false
	}

	m := map[rune]int{}

	for _, r := range magazine {
		m[r]++
	}

	for _, r := range ransomNote {
		m[r]--
		if m[r] < 0 {
			return false
		}
	}

	return true
}

func canConstruct2(ransomNote string, magazine string) bool {

	if len(ransomNote) > len(magazine) {
		return false
	}

	counter := [26]int{}

	for _, r := range magazine {
		counter[r-'a']++
	}

	for _, r := range ransomNote {
		counter[r-'a']--
		if counter[r-'a'] < 0 {
			return false
		}
	}

	return true
}

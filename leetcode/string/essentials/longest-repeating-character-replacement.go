package essentials

// https://leetcode.com/problems/longest-repeating-character-replacement/

func CharacterReplacement(s string, k int) int {

	start := 0

	// Character count in window
	counts := make([]int, 26)

	// Count of the highest occurrence of unique characters in window
	uniqueCount := 0
	maxRes := 0

	for end := 0; end < len(s); end++ {

		// Count each character occurrence in the window
		counts[s[end]-'A']++

		// Check if the recently added character has the maximum count
		uniqueCount = max(uniqueCount, counts[s[end]-'A'])

		// Number of replacement needed for all characters in window to be the same
		replaceCount := end - start + 1 - uniqueCount

		// When end-start+1-uniqueCount == 0, the window is filled with only one character
		// When end-start+1-uniqueCount > 0, then we have characters in the window that are NOT the character that occurs the most
		// end-start+1-uniqueCount is equal to exactly the # of characters that are required for all characters in window to be the same
		// Example: For a window "xxxyz", end-start+1-maxCount would equal 2. (maxCount is 3 and there are 2 characters here, "y" and "z" that are not "x" in the window.)

		if replaceCount > k {
			counts[s[start]-'A']--
			start++
		} else {
			maxRes = max(maxRes, end-start+1)
		}
	}

	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

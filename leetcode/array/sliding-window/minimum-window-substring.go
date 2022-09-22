package sliding_window

import (
	"math"
)

//https://leetcode.com/problems/minimum-window-substring/

func MinWindow(s string, t string) string {

	if len(t) > len(s) {
		return ""
	}

	// Dictionary which keeps count all the unique characters in t
	dictT := map[byte]int{}
	for _, c := range []byte(t) {
		dictT[c]++
	}

	// Number of unique characters in t, which need to be present in the desired window
	required := len(dictT)

	// Left and right pointer
	l, r := 0, 0

	// formed is used to keep track of how many unique characters in t
	// are present in the current window in its desired frequency.
	// e.g. if t is "AABC" then the window must have two A's, one B and one C.
	// Thus formed would be = 3 when all these conditions are met.
	formed := 0

	// Dictionary which keeps a count of all the unique characters in the current window.
	windowsCount := map[byte]int{}

	ans := []int{-1, 0, 0}

	for r < len(s) {

		// Add one character from the right to the window
		windowsCount[s[r]]++

		// If the frequency of the current character added equals to the
		// desired count in t then increment the formed count by 1.
		if val, ok := dictT[s[r]]; ok && windowsCount[s[r]] == val {
			formed++
		}

		// Try and contract the window till the point where it ceases to be 'desirable'.
		for formed == required {

			if ans[0] == -1 || r-l+1 < ans[0] {
				ans[0] = r - l + 1
				ans[1] = l
				ans[2] = r
			}

			// The character at the position pointed by the
			// `Left` pointer is no longer a part of the window.
			windowsCount[s[l]]--
			if val, ok := dictT[s[l]]; ok && windowsCount[s[l]] < val {
				formed--
			}

			// Move the left pointer ahead, this would help to look for a new window.
			l++
		}

		// Keep expanding the window once we are done contracting.
		r++
	}

	if ans[0] == -1 {
		return ""
	}

	return s[ans[1] : ans[2]+1]
}

func MinWindow2(s string, t string) string {

	// Create and initialize a hashmap to store the characters to be found in s, with their counts
	//(K, V) = (Character, Counts of characters in t)
	/*
		For e.g, "ABB",
		{
			A: 1,
			B: 2,
		}
		A and B both will equal or less than 0 if a valid window is found
	*/
	m := map[byte]int{}
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}

	// Left and right pointer
	// Head to save the left of min distance
	l, r, head := 0, 0, 0

	// Counter represent the number of chars of t to be found in s
	// Number of unique characters in t, which need to be present in the desired window
	counter := len(t)

	// The length of substring
	minLen := math.MaxInt

	for r < len(s) {

		// If char in s exists in t and the char count is still more than 0, decrease counter
		if m[s[r]] > 0 {
			counter--
		}

		// Decrease m[s[r]]. If char in t is encountered, m[s[r]] will be negative.
		/*
			AABBB
			012
			AB
			Counter will be zero at position 2
			The map will be in state
			{
				A: -1,
				B: 0,
			}
		*/
		m[s[r]]--
		r++

		// When we found a valid window, move start to find smaller window.
		for counter == 0 {

			// Update minimum length
			if r-l < minLen {
				head = l
				minLen = r - l
			}

			// Return the counts as the window become smaller
			m[s[l]]++

			// When char exists in t, increase counter
			if m[s[l]] > 0 {
				counter++
			}

			l++
		}
	}

	if minLen != math.MaxInt {
		return s[head : head+minLen]
	}

	return ""
}

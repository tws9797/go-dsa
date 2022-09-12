package leetcode

func lengthOfLongestSubstring(s string) int {

	hashTable := map[rune]struct{}{}
	longest := 0
	counter := 0

	for _, c := range s {

		if _, ok := hashTable[c]; !ok {
			hashTable[c] = struct{}{}
			counter++
			if counter > longest {
				longest = counter
			}
		} else {
			delete(hashTable, c)
			counter = 1
		}
	}

	return longest
}

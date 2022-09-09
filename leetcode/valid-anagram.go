package leetcode

func IsAnagram(s string, t string) bool {

	hashTable := map[rune]int{}
	if len(s) != len(t) {
		return false
	}

	for _, c := range s {
		if _, ok := hashTable[c]; !ok {
			hashTable[c] = 1
		} else {
			hashTable[c]++
		}
	}

	for _, c := range t {
		if count, ok := hashTable[c]; !ok {
			return false
		} else {
			if count == 0 {
				return false
			} else {
				hashTable[c]--
			}
		}
	}

	return true
}

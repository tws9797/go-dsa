package palindrome

import "sort"

//https://leetcode.com/problems/valid-anagram/

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

func IsAnagram2(s string, t string) bool {

	r1 := []byte(s)
	r2 := []byte(t)

	sort.Slice(r1, func(i, j int) bool {
		return r1[i] < r1[j]
	})

	sort.Slice(r2, func(i, j int) bool {
		return r2[i] < r2[j]
	})

	return string(r1) == string(r2)
}

package essentials

import (
	"sort"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/group-anagrams/

func GroupAnagrams(strs []string) [][]string {

	if len(strs) == 0 {
		return nil
	}

	m := map[string][]string{}
	var ans [][]string

	for i := 0; i < len(strs); i++ {
		str := strings.Split(strs[i], "")
		sort.Strings(str)
		anagram := strings.Join(str, "")

		m[anagram] = append(m[anagram], strs[i])
	}

	for _, v := range m {
		ans = append(ans, v)
	}

	return ans
}

func GroupAnagrams2(strs []string) [][]string {

	if len(strs) == 0 {
		return nil
	}

	groups := map[string][]string{}
	for _, s := range strs {

		freqT := [26]int{}
		for _, l := range s {
			freqT[l-'a']++
		}

		hashS := ""
		for _, f := range freqT {
			hashS = hashS + strconv.Itoa(f) + ","
		}
		hashS = hashS[:len(hashS)-1]

		groups[hashS] = append(groups[hashS], s)
	}

	// convert from hash map to 2d array
	var result [][]string
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

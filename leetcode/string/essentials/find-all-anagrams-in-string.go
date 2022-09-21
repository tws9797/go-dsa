package essentials

import "reflect"

func FindAnagrams(s string, p string) []int {

	m := map[byte]int{}
	for i := 0; i < len(p); i++ {
		m[p[i]]++
	}
	count := len(p)
	var ans []int
	l, r := 0, 0

	for r < len(s) {

		if m[s[r]] > 0 {
			count--
		}
		m[s[r]]--

		for count == 0 {

			m[s[l]]++
			if m[s[l]] > 0 {
				if r-l+1 == len(p) {
					ans = append(ans, l)
				}

				count++
			}
			l++
		}

		r++
	}

	return ans
}

func FindAnagrams2(s string, p string) []int {
	ns := len(s)
	np := len(p)

	if ns < np {
		return nil
	}

	pCount := map[byte]int{}
	sCount := map[byte]int{}

	// Build reference hashmap using string p
	for _, ch := range []byte(p) {
		pCount[ch]++
	}

	var output []int

	// Sliding window on the string s
	for i := 0; i < ns; i++ {

		// add one more letter on the right side of the window
		sCount[s[i]]++

		// remove one letter from the left side of the window
		//cbaabcbacd
		//abc
		/*
			b:1
			a:2
		*/

		if i >= np {
			if sCount[s[i-np]] == 1 {
				delete(sCount, s[i-np])
			} else {
				sCount[s[i-np]]--
			}
		}

		// compare hashmap in the sliding window
		// with the reference hashmap
		if reflect.DeepEqual(pCount, sCount) {
			output = append(output, i-np+1)
		}
	}

	return output
}

func FindAnagrams3(s string, p string) []int {
	ns := len(s)
	np := len(p)

	if ns < np {
		return nil
	}

	pCount := [26]int{}
	sCount := [26]int{}

	// Build reference hashmap using string p
	for _, ch := range p {
		pCount[ch-'a']++
	}

	var output []int

	// Sliding window on the string s
	for i := 0; i < ns; i++ {

		// add one more letter on the right side of the window
		sCount[s[i]-'a']++

		if i >= np {
			sCount[s[i-np]-'a']--
		}

		// compare hashmap in the sliding window
		// with the reference hashmap
		if pCount == sCount {
			output = append(output, i-np+1)
		}
	}

	return output
}

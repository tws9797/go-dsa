package essentials

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

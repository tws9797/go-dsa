package two_pointers

func CountSubstrings(s string) int {
	count := 0

	for i := 0; i < len(s); i++ {
		count += extendPalindrome(s, i, i)
		count += extendPalindrome(s, i, i+1)
	}

	return count
}

func extendPalindrome(s string, l, r int) int {

	res := 0
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
		res++
	}

	return res
}

func CountSubstringsDP(s string) int {

	if len(s) == 0 {
		return 0
	}

	count, n := 0, len(s)
	d := make([][]bool, n)
	for i := 0; i < n; i++ {
		d[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				d[i][j] = true
			} else if i+1 == j {
				d[i][j] = s[i] == s[j]
			} else {
				d[i][j] = s[i] == s[j] && d[i+1][j-1]
			}

			if d[i][j] {
				count++
			}
		}
	}

	return count
}

package essentials

// https://leetcode.com/problems/roman-to-integer/

func RomanToInt(s string) int {

	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var sum int

	for i := len(s) - 1; i >= 0; i-- {

		if (sum > 3 && s[i] == 'I') || (sum > 40 && s[i] == 'X') || (sum > 400 && s[i] == 'C') {
			sum -= m[s[i]]
		} else {
			sum += m[s[i]]
		}
	}

	return sum
}

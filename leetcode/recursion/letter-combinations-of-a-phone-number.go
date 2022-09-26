package recursion

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

func LetterCombinations(digits string) []string {

	if len(digits) == 0 {
		return nil
	}

	m := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	var ans []string

	strs := []byte(digits)

	backtrackLetterCombinations(0, nil, strs, m, &ans)

	return ans
}

func backtrackLetterCombinations(start int, prev []byte, strs []byte, m map[byte][]byte, ans *[]string) {

	if len(prev) == len(strs) {
		*ans = append(*ans, string(prev))
		return
	}

	for _, v := range m[strs[start]] {
		prev = append(prev, v)
		backtrackLetterCombinations(start+1, prev, strs, m, ans)
		prev = prev[:len(prev)-1]
	}
}

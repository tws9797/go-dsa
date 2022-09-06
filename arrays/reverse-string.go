package arrays

// https://leetcode.com/problems/reverse-string/
func ReverseString(s []byte) {

	if len(s) > 1 {
		for low, high := 0, len(s)-1; low <= high; low, high = low+1, high-1 {
			s[low], s[high] = s[high], s[low]
		}
	}
}

package recursion

// https://leetcode.com/problems/reverse-string/

func ReverseString(s []byte) {

	if len(s) > 1 {
		for low, high := 0, len(s)-1; low < high; low, high = low+1, high-1 {
			s[low], s[high] = s[high], s[low]
		}
	}
}

func ReverseStringRecursive(s []byte) {
	helper(s, 0, len(s)-1)
}

func helper(s []byte, left int, right int) {
	if left >= right {
		return
	}

	s[left], s[right] = s[right], s[left]
	left++
	right--
	helper(s, left, right)
}

package leetcode

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	s = strings.ToUpper(s)
	l, r := 0, len(s)-1

	for l < r {
		if !unicode.IsLetter(rune(s[l])) && !unicode.IsNumber(rune(s[l])) {
			l++
			continue
		}
		if !unicode.IsLetter(rune(s[r])) && !unicode.IsNumber(rune(s[r])) {
			r--
			continue
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

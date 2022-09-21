package palindrome

import (
	"strings"
	"unicode"
)

// https://leetcode.com/problems/valid-palindrome/

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

func IsPalindrome2(s string) bool {

	s = strings.ToUpper(s)
	l, r := 0, len(s)-1

	for l < r {
		if !checkAlphanumeric(s[l]) {
			l++
			continue
		}

		if !checkAlphanumeric(s[r]) {
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

func checkAlphanumeric(b byte) bool {
	return ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9')
}

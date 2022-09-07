package recursion

func ReverseString(s string) string {

	lastCharacterIdx := len(s) - 1

	if len(s) == 1 {
		return string(s[lastCharacterIdx])
	}

	return string(s[lastCharacterIdx]) + ReverseString(s[:lastCharacterIdx])
}

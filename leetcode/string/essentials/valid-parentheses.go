package essentials

// https://leetcode.com/problems/valid-parentheses/

func IsValid(s string) bool {
	stack := make([]byte, 0)
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, c := range []byte(s) {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != m[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

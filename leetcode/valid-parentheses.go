package leetcode

func IsValidStack(s string) bool {

	var stackR []rune

	for _, r := range s {

		if r == '(' || r == '[' || r == '{' {
			stackR = append(stackR, r)
		} else {
			// To prevent scenario where no r is append. For e.g, ]] ))
			if len(stackR) > 0 {
				openBracket := stackR[len(stackR)-1]
				if (openBracket == '{' && r != '}') || (openBracket == '[' && r != ']') || (openBracket == '(' && r != ')') {
					return false
				} else {
					stackR = stackR[:len(stackR)-1]
				}
			} else {
				return false
			}
		}
	}

	// To prevent scenario where the else case is not triggered. For e.g, ((
	if len(stackR) > 0 {
		return false
	}

	return true
}

func IsValidIterative(str string) bool {

	var s []byte

	if str[0] == '(' || str[0] == '{' || str[0] == '[' {

		for i := 0; i < len(str); i++ {
			s = append(s, str[i])
			if len(s) > 1 {
				if s[len(s)-2] == '[' && s[len(s)-1] == ']' {
					s = s[:len(s)-2]
				} else if s[len(s)-2] == '{' && s[len(s)-1] == '}' {
					s = s[:len(s)-2]
				} else if s[len(s)-2] == '(' && s[len(s)-1] == ')' {
					s = s[:len(s)-2]
				}
			}
		}

		if len(s) == 0 {
			return true
		}
	}

	return false
}

func IsValidStack2(s string) bool {

	stack := make([]rune, 0)
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':

			// If no stack during the loop, means the number of open bracket does not match that of closing bracket
			// If pop of stack does not match
			if len(stack) == 0 || stack[len(stack)-1] != m[c] {
				return false
			}

			// Reduce length after pop a stack
			stack = stack[:len(stack)-1]
		}
	}

	// If there is still stack left after the loop,  means the number of open bracket does not match that of closing bracket
	return len(stack) == 0
}

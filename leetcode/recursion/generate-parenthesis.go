package recursion

// https://leetcode.com/problems/generate-parentheses/

func GenerateParenthesis(n int) []string {

	m := map[int][]string{}

	return helperGenerateParenthesis(n, m)
}

func helperGenerateParenthesis(n int, m map[int][]string) []string {

	if ans, ok := m[n]; ok {
		return ans
	}

	var ans []string
	if n == 0 {
		ans = append(ans, "")
	} else {
		for i := 0; i < n; i++ {
			for _, left := range helperGenerateParenthesis(i, m) {
				for _, right := range helperGenerateParenthesis(n-1-i, m) {
					ans = append(ans, "("+left+")"+right)
				}
			}
		}
	}

	m[n] = ans

	return ans
}

func GenerateParenthesisBacktracking(n int) []string {
	var ans []string
	var str string
	backtrack(&ans, str, 0, 0, n)
	return ans
}

func backtrack(ans *[]string, str string, open int, close int, n int) {

	if len(str) == n*2 {
		*ans = append(*ans, str)
		return
	}

	// Only add open parenthesis when open is less than n
	if open < n {
		str += "("
		backtrack(ans, str, open+1, close, n)
		str = str[:len(str)-1]
	}

	// Only add closing parenthesis when open parenthesis is more than close
	if close < open {
		str += ")"
		backtrack(ans, str, open, close+1, n)
		str = str[:len(str)-1]
	}
}

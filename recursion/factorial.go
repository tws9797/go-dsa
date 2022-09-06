package recursion

func FindFactorialRecursive(num int) int {

	if num == 0 {
		return 1
	}

	if num == 1 {
		return num
	}

	return num * FindFactorialRecursive(num-1)
}

func FindFactorialIterative(num int) int {

	res := 1
	for i := num; i > 0; i-- {
		res *= i
	}

	return res
}

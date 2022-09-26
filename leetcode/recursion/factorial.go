package recursion

func FactorialIterative(num int) int {

	res := 1

	if num >= 1 {
		res *= num
		num--
	}

	return res
}

func FactorialRecursive(num int) int {
	if num == 0 || num == 1 {
		return num
	}

	return num * FactorialRecursive(num-1)
}

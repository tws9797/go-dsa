package recursion

func FibonacciIterative(n int) int {

	firstNum := 0
	secondNum := 1

	//sum 0
	//lastSum 0
	for i := 0; i < n-1; i++ {
		temp := secondNum
		secondNum = firstNum + secondNum
		firstNum = temp
	}

	return secondNum
}

func FibonacciRecursive(n int) int {

	if n < 2 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

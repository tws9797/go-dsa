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

func FibonacciLoop(n int) int {

	fibBox := []int{0, 1}

	//0 1 2 3 4 5 6 7
	//0 1 1 2 3 5 8 13

	//0 1  0
	//1 1  1
	//1 2  2
	//2 3  3
	//3 5  4
	//5 8  5
	//8 13 6

	for i := 0; i < n; i++ {
		v := fibBox[i] + fibBox[i+1]
		fibBox = append(fibBox, v)
	}

	result := fibBox[n]

	return result
}

func FibonacciMemoize(n int, cache map[int]int) int {

	if n < 2 {
		cache[n] = n
		return n
	}

	if _, ok := cache[n-1]; !ok {
		cache[n-1] = FibonacciMemoize(n-1, cache)
	}

	if _, ok := cache[n-2]; !ok {
		cache[n-2] = FibonacciMemoize(n-2, cache)
	}

	return cache[n-1] + cache[n-2]
}

func FibonacciMemoizeMaster(n int) int {
	cache := make(map[int]int)
	bucket := make([]int, n)

	for i := 0; i < n; i++ {
		//1 2 3 4 5 6 7
		//0 1 2 3 4 5 6
		//1 1
		bucket[i] = FibonacciMemoize(i+1, cache)
	}

	result := bucket[n-1]
	return result
}

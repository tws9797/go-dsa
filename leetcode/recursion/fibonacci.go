package recursion

func FibonacciIterative(num int) int {
	first := 0
	second := 1

	for i := 0; i < num; i++ {
		temp := first
		first = second
		second = temp + first
	}

	return first
}

func FibonacciRecursive(num int) int {
	if num <= 1 {
		return num
	}

	return FibonacciRecursive(num-1) * FibonacciRecursive(num-2)
}

func FibonacciMemoizeMaster(num int) int {
	cache := make(map[int]int)
	bucket := make([]int, num)

	if num == 0 {
		return 0
	}

	for i := 0; i < num; i++ {
		bucket[i] = FibonacciMemoize(i+1, cache)
	}

	return bucket[num-1]
}

func FibonacciMemoize(num int, cache map[int]int) int {

	if num < 2 {
		cache[num] = num
		return num
	}

	if _, ok := cache[num-1]; !ok {
		cache[num-1] = FibonacciMemoize(num-1, cache)
	}

	if _, ok := cache[num-2]; !ok {
		cache[num-2] = FibonacciMemoize(num-2, cache)
	}

	return cache[num-1] + cache[num-2]
}

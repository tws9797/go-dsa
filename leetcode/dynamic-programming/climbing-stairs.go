package dynamic_programming

func ClimbStairsBF(n int) int {

	return ClimbStairsRecursive(0, n)
}

func ClimbStairsRecursive(start int, n int) int {

	if start > n {
		return 0
	}

	if start == n {
		return 1
	}

	return ClimbStairsRecursive(start+1, n) + ClimbStairsRecursive(start+2, n)
}

func ClimbStairsBFMemoization(n int) int {

	memo := make([]int, n)
	return ClimbStairsRecursiveMemoization(0, n, memo)
}

func ClimbStairsRecursiveMemoization(start int, n int, memo []int) int {

	if start > n {
		return 0
	}

	if start == n {
		return 1
	}

	if memo[start] > 0 {
		return memo[start]
	}

	memo[start] = ClimbStairsRecursiveMemoization(start+1, n, memo) + ClimbStairsRecursiveMemoization(start+2, n, memo)
	return memo[start]
}

func ClimbStairsDP(n int) int {

	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func ClimbStairsFibonacciNumber(n int) int {

	if n <= 2 {
		return n
	}

	ways := 0
	twoBelow := 1
	oneBelow := 2
	for i := 3; i <= n; i++ {
		ways = oneBelow + twoBelow
		twoBelow = oneBelow
		oneBelow = ways
	}

	return ways
}

package dynamic_programming

import "math"

func CoinChange(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}
	//count := make([]int, amount+1)
	return coinChangeRecursive(coins, amount)
}

func coinChangeRecursive(coins []int, remain int) int {

	if remain < 0 {
		return -1
	}

	if remain == 0 {
		return 0
	}

	min := math.MaxInt
	for _, coin := range coins {
		count := coinChangeRecursive(coins, remain-coin)
		if count == -1 {
			continue
		}
		if count+1 < min {
			min = count + 1
		}
	}

	if min == math.MaxInt {
		return -1
	}

	return min
}

func coinChangeTopdown(coins []int, remain int, count []int) int {
	if remain < 0 {
		return -1
	}
	if remain == 0 {
		return 0
	}
	if count[remain] != 0 {
		return count[remain]
	}

	min := math.MaxInt

	for _, coin := range coins {
		res := coinChangeTopdown(coins, remain-coin, count)
		if res == -1 {
			continue
		}
		if res < min {
			min = 1 + res
		}
	}

	if min == math.MaxInt {
		count[remain] = -1
	} else {
		count[remain] = min
	}
	return count[remain]
}

func coinChangeBottomUp(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			if dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}

		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

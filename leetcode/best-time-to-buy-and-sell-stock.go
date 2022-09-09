package leetcode

import "math"

func maxProfit(prices []int) int {

	profit := 0

	// min price act as a pointer to be deducted
	minPrice := math.MaxInt

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > profit {
			profit = prices[i] - profit
		}
	}

	return profit
}

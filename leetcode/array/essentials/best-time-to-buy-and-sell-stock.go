package essentials

import "math"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func MaxProfit(prices []int) int {

	maxProfit := 0

	// min price act as a pointer to be deducted
	minPrice := math.MaxInt

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}

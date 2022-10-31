package essentials

import "math"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
/*
	Input: prices = [7,1,5,3,6,4]
	Output: 5
	Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
	Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

	Input: prices = [7,6,4,3,1]
	Output: 0
	Explanation: In this case, no transactions are done and the max profit = 0.
*/

func MaxProfit(prices []int) int {

	// Maintain two pointer, the lowest price and maximum profit
	maxProfit := 0
	minPrice := math.MaxInt

	for i := 0; i < len(prices); i++ {

		// If current price is smaller than the minPrice, set the current price to minPrice
		// Note: Only use this loop to set the minPrice, since there is no prices lower than current price
		// No point of check the maxProfit
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}

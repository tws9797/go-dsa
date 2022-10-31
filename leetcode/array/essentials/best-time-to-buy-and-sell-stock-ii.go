package essentials

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
/*
	Input: prices = [7,1,5,3,6,4]
	Output: 7
	Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
	Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
	Total profit is 4 + 3 = 7.

	Input: prices = [1,2,3,4,5]
	Output: 4
	Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
	Total profit is 4.

	Input: prices = [7,6,4,3,1]
	Output: 0
	Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.
*/

func MaxProfit2(prices []int) int {

	totalProfit := 0

	for i := 1; i < len(prices); i++ {

		// Cumulative profit from a to c (a,b,c), is equal to (b-a) + (c-b)
		// Keep on adding the difference between the consecutive numbers of the array if the second number is larger than the first one,
		// At the total sum we obtain will be the maximum profit.
		if prices[i] > prices[i-1] {
			totalProfit += prices[i] - prices[i-1]
		}
	}

	return totalProfit
}

package essentials

// https://leetcode.com/problems/maximum-product-subarray/

func MaxProduct(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	// Keep track of the accumulated product of positive number
	// This value will be picked if the accumulated product has been steadily increasing (all positive numbers).
	maxSoFar := nums[0]

	// Single negative number can flip the combo chain
	// Handle negative number
	// This value will be picked if the current number is a negative number and the combo chain has been disrupted by a single negative number before
	//(In a sense, this value is like an antidote to an already poisoned combo chain).
	minSoFar := nums[0]

	result := maxSoFar

	for i := 1; i < len(nums); i++ {

		/*
			This value will be picked if the accumulated product has been really bad (even compared to the current number).
			Combo chain is broken
			This can happen when the current number has a preceding zero (e.g. [0,4]) or is preceded by a single negative number (e.g. [-3,5]).
		*/
		curr := nums[i]
		tempMax := max(curr, max(maxSoFar*curr, minSoFar*curr))
		minSoFar = min(curr, min(maxSoFar*curr, minSoFar*curr))

		maxSoFar = tempMax

		result = max(maxSoFar, result)
	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

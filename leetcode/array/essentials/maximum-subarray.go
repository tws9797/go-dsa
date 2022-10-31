package essentials

import "math"

// https://leetcode.com/problems/maximum-subarray/

func maxSubArray(nums []int) int {

	sum := 0
	maxSum := math.MinInt

	for _, num := range nums {

		sum += num

		if sum > maxSum {
			maxSum = sum
		}

		// Any subarray whose sum is positive is worth keeping
		// Have possibility of become greater value
		// Whenever the sum of the array is negative,
		// The entire array is not worth keeping, reset it back to an empty array.
		if sum < 0 {
			sum = 0
		}

	}

	return maxSum

}

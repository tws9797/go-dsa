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

		if sum < 0 {
			sum = 0
		}

	}

	return maxSum

}

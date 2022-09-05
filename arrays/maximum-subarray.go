package arrays

import "math"

func maxSubArray(nums []int) int {

	sum := 0
	max := math.MinInt

	//If sum of subarray is positive, it has possible to make the next value bigger
	// so keep do it until it turn to negative
	//If the sum is negative, it has no use to the next element, so break
	for _, val := range nums {

		sum += val

		if sum > max {
			max = sum
		}

		if sum < 0 {
			sum = 0
		}
	}

	return max
}

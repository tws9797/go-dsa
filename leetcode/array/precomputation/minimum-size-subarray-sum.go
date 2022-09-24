package precomputation

import "math"

// https://leetcode.com/problems/minimum-size-subarray-sum/

func MinSubArrayLen(target int, nums []int) int {

	l, r, sum := 0, 0, 0
	minLen := math.MaxInt

	for r < len(nums) {

		if nums[r] == target {
			return 1
		}

		sum += nums[r]
		r++

		for sum >= target {

			if r-l < minLen {
				minLen = r - l
			}

			sum -= nums[l]

			l++
		}

	}

	if minLen == math.MaxInt {
		return 0
	}

	return minLen
}

package leetcode

func MaxSubArray(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	currentSubArray := nums[0]
	maxSubArray := nums[0]

	//If sum of subarray is positive, it has possible to make the next value bigger
	// so keep do it until it turn to negative
	//If the sum is negative, it has no use to the next element, so break
	for _, num := range nums[1:] {

		if currentSubArray+num > num {
			currentSubArray += num
		} else {
			currentSubArray = num
		}

		if maxSubArray < currentSubArray {
			maxSubArray = currentSubArray
		}
	}

	return maxSubArray
}

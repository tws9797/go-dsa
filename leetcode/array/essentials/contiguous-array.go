package essentials

// https://leetcode.com/problems/contiguous-array/
/*
	Input: nums = [0,1]
	Output: 2
	Explanation: [0, 1] is the longest contiguous subarray with an equal number of 0 and 1.

	Input: nums = [0,1,0]
	Output: 2
	Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
*/

func FindMaxLength(nums []int) int {

	m := map[int]int{}

	// Initialize counter 0 for case like [1,0]
	m[0] = -1
	maxLen := 0
	count := 0

	for i := 0; i < len(nums); i++ {

		if nums[i] == 1 {
			count++
		} else {
			count--
		}

		// Same count indicates the number of zero is the same between the indices of i and val
		if val, ok := m[count]; ok {
			maxLen = max(maxLen, i-val)
		} else {
			m[count] = i
		}
	}

	return maxLen
}

package essentials

// https://leetcode.com/problems/contiguous-array/

func FindMaxLength(nums []int) int {

	m := map[int]int{}
	m[0] = -1
	maxLen := 0
	count := 0

	for i := 0; i < len(nums); i++ {

		if nums[i] == 1 {
			count++
		} else {
			count--
		}

		if val, ok := m[count]; ok {
			maxLen = max(maxLen, i-val)
		} else {
			m[count] = i
		}
	}

	return maxLen
}

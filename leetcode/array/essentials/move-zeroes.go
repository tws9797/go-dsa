package essentials

// https://leetcode.com/problems/move-zeroes/

func MoveZeroes(nums []int) {

	for i, count := 0, 0; i < len(nums); i++ {

		// Use count as pointer to the non-zeros
		if nums[i] != 0 {
			nums[i], nums[count] = 0, nums[i]
			count++
		}
	}
}

func moveZeroes2(nums []int) {

	lastNonZeroFoundAt := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZeroFoundAt] = nums[i]
			lastNonZeroFoundAt++
		}
	}

	for i := lastNonZeroFoundAt; i < len(nums); i++ {
		nums[i] = 0
	}
}

package two_pointers

// https://leetcode.com/problems/sort-colors/
/*
	Input: nums = [2,0,2,1,1,0]
	Output: [0,0,1,1,2,2]
*/

func SortColors(nums []int) {

	// Three pointers to track the rightmost boundary of zeroes, leftmost boundary of twos and the current element
	l := 0
	curr := 0
	r := len(nums) - 1

	for curr <= r {

		// Increase curr if the current element is in the right place
		if nums[curr] == 0 {
			nums[curr], nums[l] = nums[l], nums[curr]
			l++
			curr++
		} else if nums[curr] == 2 {

			// Put 2 in the right place, no need to increase curr as nums[curr] might not in the right place
			nums[curr], nums[r] = nums[r], nums[curr]
			r--
		} else {

			// 1 will be leave as it is
			curr++
		}
	}
}

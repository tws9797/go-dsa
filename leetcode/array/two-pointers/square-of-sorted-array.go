package two_pointers

// https://leetcode.com/problems/squares-of-a-sorted-array/

func SortedSquares(nums []int) []int {

	l, r := 0, len(nums)-1
	ans := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {

		if nums[l]*nums[l] >= nums[r]*nums[r] {
			ans[i] = nums[l] * nums[l]
			l++
		} else {
			ans[i] = nums[r] * nums[r]
			r--
		}
	}

	return ans
}

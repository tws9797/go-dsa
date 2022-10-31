package essentials

// https://leetcode.com/problems/binary-search/

func Search(nums []int, target int) int {

	l, r := 0, len(nums)-1

	for l <= r {

		// https://stackoverflow.com/questions/27167943/why-leftright-left-2-will-not-overflow
		// Prevent r too big and causes overflow
		mid := r + (l-r)/2

		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

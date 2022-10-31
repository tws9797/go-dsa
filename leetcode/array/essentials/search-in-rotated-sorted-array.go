package essentials

// https://leetcode.com/problems/search-in-rotated-sorted-array/

func SearchRotated(nums []int, target int) int {

	start, end := 0, len(nums)-1

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[start] {

			// For the array like [3,1]
			// Pivot element is larger than the first element in the array
			// i.e. the subarray from the first element to the pivot is non-rotated
			// If the target is located in the non-rotated subarray:
			// go left: `end = mid - 1`.
			// Otherwise: go right: `start = mid + 1`.
			if target >= nums[start] && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {

			// Pivot element is smaller than the first element of the array
			// i.e. the rotation index is somewhere between 0 and mid.
			// It implies that the sub-array from the pivot element to the last one is non-rotated:
			// If the target is located in the non-rotated subarray:
			// go right: `start = mid + 1`.
			// Otherwise: go left: `end = mid - 1`.
			if target <= nums[end] && target > nums[mid] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}

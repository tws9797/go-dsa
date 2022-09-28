package essentials

// https://leetcode.com/problems/rotate-array/

func Rotate(nums []int, k int) {

	lenNums := len(nums)

	if lenNums > 1 {

		res := make([]int, lenNums)

		// [1,2,3,4,5]
		// [0,0,0,0,0]
		// [0,0,1,0,0]
		// [0,0,1,2,0]
		// [0,0,1,2,3]
		// [4,0,1,2,3]
		// [4,5,1,2,3]

		// Put the number in the right position
		for i := 0; i < lenNums; i++ {
			res[(i+k)%lenNums] = nums[i]
		}

		// Copy res to nums
		for i := 0; i < lenNums; i++ {
			nums[i] = res[i]
		}
	}
}

func Rotate2(nums []int, k int) {

	if k == 0 {
		return
	}

	lenNums := len(nums)
	count := 0

	for i := 0; count < lenNums; i++ {

		idxMoveTo := i

		// Prepare a temporary placeholder for the replaced element
		prev := nums[i]

		for ok := true; ok; ok = idxMoveTo != i {

			// Get the next index to move the pointed element
			idxMoveTo = (idxMoveTo + k) % lenNums

			// Swap the temporary placeholder with the pointed element
			nums[idxMoveTo], prev = prev, nums[idxMoveTo]

			// Increase the count for each swap
			// By theory the count for each swap is smaller than the length of nums array
			count++
		}
	}

}

func RotateIterative(nums []int, k int) {

	k %= len(nums)

	// Move the array to right for k number of times
	for i := 0; i < k; i++ {

		prev := nums[len(nums)-1]

		for j := 0; j < len(nums); j++ {
			temp := nums[j]
			nums[j] = prev
			prev = temp
		}
	}
}

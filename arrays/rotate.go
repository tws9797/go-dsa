package arrays

// https://leetcode.com/problems/rotate-array/
func Rotate(nums []int, k int) {

	lenNums := len(nums)

	if lenNums > 1 {

		res := make([]int, lenNums)

		for i := 0; i < lenNums; i++ {
			idxMoveTo := i + k
			if idxMoveTo >= lenNums {
				idxMoveTo %= lenNums
			}

			res[idxMoveTo] = nums[i]
		}

		for i := 0; i < lenNums; i++ {
			nums[i] = res[i]
		}
	}
}

func Rotate2(nums []int, k int) {

	lenNums := len(nums)

	if k == 0 {
		return
	}

	if nums == nil || lenNums == 0 {
		return
	}

	if lenNums > 1 {

		count := 0

		for i := 0; count < lenNums; i++ {

			// Prepare a temporary placeholder for the replaced element
			tmpNum := nums[i]

			// Get the next index to move the pointed element
			idxMoveTo := (i + k) % lenNums

			// For the even number of k
			for idxMoveTo != i {

				// Swap the temporary placeholder with the pointed element
				nums[idxMoveTo], tmpNum = tmpNum, nums[idxMoveTo]
				idxMoveTo = (idxMoveTo + k) % lenNums

				// Increase the count for each swap
				// By theory the count for each swap is smaller than the length of nums array
				count++
			}

			// The loop will break when idxMoveTo reach i
			// The number stored in temporary placeholder have to be executed one last time
			nums[idxMoveTo], tmpNum = tmpNum, nums[idxMoveTo]
			count++
		}

	}
}

package essentials

// https://leetcode.com/problems/sliding-window-maximum/

func MaxSlidingWindow(nums []int, k int) []int {

	ans := make([]int, len(nums)-k+1)
	prevHigh := 0

	for i := 0; i < k; i++ {
		if nums[i] >= nums[prevHigh] {
			prevHigh = i
		}
	}

	ans[0] = nums[prevHigh]

	//[1,3,1,2,0,5]

	for i := 1; i < len(nums)-k+1; i++ {

		// Index of last element in the window
		lastElWindowIdx := i + k - 1

		// If index for highest of previous window is in the range of current window
		if nums[lastElWindowIdx] >= nums[prevHigh] {

			// Compare the last element to the previous window highest
			// If equal or highest, set the prevHigh to the index of last element of current window
			prevHigh = lastElWindowIdx
		} else if prevHigh < i {

			highest := nums[i]
			prevHigh = i

			// If index for highest of previous window is not in the range of current window
			// That is, i-1
			for j := i; j < i+k; j++ {
				if nums[j] >= highest {
					highest = nums[j]
					prevHigh = j
				}
			}
		}

		ans[i] = nums[prevHigh]
	}

	return ans
}

func MaxSlidingWindowDP(nums []int, k int) []int {

	n := len(nums)
	if n*k == 0 {
		return nil
	}

	if k == 1 {
		return nums
	}

	left := make([]int, n)
	left[0] = nums[0]

	right := make([]int, n)
	right[n-1] = nums[n-1]

	for i := 1; i < n; i++ {

		// Iterate along the array in the direction left->right and build an array left
		if i%k == 0 {
			left[i] = nums[i]
		} else {
			if left[i-1] >= nums[i] {
				left[i] = left[i-1]
			} else {
				left[i] = nums[i]
			}
		}

		// Iterate along the array in the direction right->left and build an array right
		j := n - i - 1
		if (j+1)%k == 0 {
			right[j] = nums[j]
		} else {
			if right[j+1] >= nums[j] {
				right[j] = right[j+1]
			} else {
				right[j] = nums[j]
			}
		}

	}

	output := make([]int, n-k+1)

	// Build an output array as max(right[i], left[i + k - 1]) for i in range (0, n - k + 1)
	for i := 0; i < n-k+1; i++ {
		if left[i+k-1] >= right[i] {
			output[i] = left[i+k-1]
		} else {
			output[i] = right[i]
		}
	}

	return output
}

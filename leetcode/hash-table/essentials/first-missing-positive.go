package essentials

// https://leetcode.com/problems/first-missing-positive/

func FirstMissingPositive(nums []int) int {

	curr := 0
	lenNums := len(nums)

	for curr < lenNums {

		// Move current number to the index it meant to be
		idxToMoveTo := nums[curr] - 1

		// If the idxToMoveTo is not out of bound
		/*
			For e.g, number 0 or number -2 definitely will result in idxToMoveTo -1 and -3 which is out of bound [0, lenNums)
		*/
		// Repeatedly change the current index number until the current index number is with the right number or out of bound number
		// Then proceed move to the next index

		// nums[curr] != nums[idxToMoveTo] is to prevent duplicate number such as [1,1]
		if 0 <= idxToMoveTo && idxToMoveTo < lenNums && nums[curr] != nums[idxToMoveTo] {
			nums[curr], nums[idxToMoveTo] = nums[idxToMoveTo], nums[curr]
		} else {
			curr++
		}
	}

	// Check if the number are in the correct position
	/*
		For e.g, index 0 should have num of 1 (val should equal i+1)
	*/
	for i, val := range nums {
		if val != i+1 {
			return i + 1
		}
	}

	// If all number are in the right position, return the length of nums + 1
	return lenNums + 1
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func FirstMissingPostivie2(nums []int) int {

	lenNums := len(nums)
	contains := 0

	for i := 0; i < lenNums; i++ {
		if nums[i] == 1 {
			contains++
			break
		}
	}

	if contains == 0 {
		return 1
	}

	for i := 0; i < lenNums; i++ {
		if nums[i] <= 0 || nums[i] > lenNums {
			nums[i] = 1
		}
	}

	for i := 0; i < lenNums; i++ {

		a := Abs(nums[i])

		if a == lenNums {
			nums[0] = -Abs(nums[0])
		} else {
			nums[a] = -Abs(nums[a])
		}
	}

	for i := 1; i < lenNums; i++ {
		if nums[i] > 0 {
			return i
		}
	}

	if nums[0] > 0 {
		return lenNums
	}

	return lenNums + 1
}

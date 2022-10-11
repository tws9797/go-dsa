package binary

// https://leetcode.com/problems/missing-number/

func MissingNumber(nums []int) int {

	curr := 0

	for curr < len(nums) {

		if nums[curr] >= len(nums) || nums[curr] == curr {
			curr++
		} else {
			nums[curr], nums[nums[curr]] = nums[nums[curr]], nums[curr]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}

	return len(nums)
}

func MissingNumberBitManipulation(nums []int) int {

	/*
		XOR - Only true when the arguments is different
		x ^ x = 0
		0 ^ x = x
		x ^ y = y ^ x
		e.g, 4,2,1,0
		missing = 4
		0100 ^= 0000 ^ 0100
		0000 ^= 0001 ^ 0010
		0011 ^= 0010 ^ 0001
		0011 ^= 0000 ^ 0011
	*/
	missing := len(nums)
	for i := 0; i < len(nums); i++ {
		missing ^= i ^ nums[i]
	}

	return missing
}

func MissingNumberGaussFormula(nums []int) int {
	expectedSum := len(nums) * (len(nums) + 1) / 2
	actualSum := 0

	for _, num := range nums {
		actualSum += num
	}

	return expectedSum - actualSum
}

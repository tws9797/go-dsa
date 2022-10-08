package binary

// https://leetcode.com/problems/single-number/

func SingleNumber(nums []int) int {

	/*
		For e.g,
		[4,2,1,2,1]
		0100->0100
		0010->0110
		0001->0111
		0010->0101
		0001->0100
		ans 0100
	*/
	a := 0

	for _, val := range nums {
		a ^= val
	}

	return a
}

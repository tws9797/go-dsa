package essentials

// https://leetcode.com/problems/house-robber/

func Rob(nums []int) int {

	memo := make([]int, 100)

	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}

	return robFrom(memo, 0, nums)
}

func robFrom(memo []int, i int, nums []int) int {

	if i >= len(nums) {
		return 0
	}

	// Return cached value.
	if memo[i] > -1 {
		return memo[i]
	}

	// Either rob from the next room, or current room + max(robFrom(i+2))
	ans := max(robFrom(memo, i+1, nums), robFrom(memo, i+2, nums)+nums[i])

	// Cache for future use.
	memo[i] = ans

	return ans
}

func RobDP(nums []int) int {
	N := len(nums)

	if N == 0 {
		return 0
	}

	// Append a 0 for first i+2 scenario
	maxRobbedAmount := make([]int, len(nums)+1)
	maxRobbedAmount[N-1] = nums[N-1]

	// robFrom(i)=max(robFrom(i+1),robFrom(i+2)+nums(i))
	for i := N - 2; i >= 0; i-- {
		maxRobbedAmount[i] = max(maxRobbedAmount[i+1], maxRobbedAmount[i+2]+nums[i])
	}

	return maxRobbedAmount[0]
}

func RobDP2(nums []int) int {

	N := len(nums)

	if N == 0 {
		return 0
	}

	robNextPlusOne := 0
	robNext := nums[N-1]

	for i := N - 2; i >= 0; i-- {
		current := max(robNext, robNextPlusOne+nums[i])

		robNextPlusOne = robNext
		robNext = current
	}

	return robNext
}

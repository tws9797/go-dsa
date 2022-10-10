package dynamic_programming

// https://leetcode.com/problems/partition-equal-subset-sum/

func CanPartitionTopdown(nums []int) bool {

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	subsetSum := sum / 2
	memo := make([][]int, len(nums)+1)

	for i := range memo {
		memo[i] = make([]int, subsetSum+1)
	}

	return canPartitionTopdown(nums, 0, subsetSum, memo)
}

func canPartitionTopdown(nums []int, index int, sum int, memo [][]int) bool {
	if sum == 0 {
		return true
	}

	if sum < 0 || index == len(nums) {
		return false
	}

	if memo[index][sum] != 0 {
		return memo[index][sum] == 1
	}

	res := canPartitionTopdown(nums, index+1, sum-nums[index], memo) || canPartitionTopdown(nums, index+1, sum, memo)
	if res {
		memo[index][sum] = 1
	} else {
		memo[index][sum] = -1
	}

	return res
}

func CanPartitionBottomUp(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	subsetSum := sum / 2
	dp := make([][]bool, len(nums)+1)

	for i := range dp {
		dp[i] = make([]bool, subsetSum+1)
	}

	/*
		dp[i][j] means whether the specific sum j can be gotten from the first i number
	*/
	dp[0][0] = true
	for i := 1; i <= len(nums); i++ {
		curr := nums[i-1]
		for j := 0; j <= subsetSum; j++ {
			if j < curr {
				dp[i][j] = dp[i-1][j]
			} else {

				// To pick nums[i] or not to pick 
				dp[i][j] = dp[i-1][j] || dp[i-1][j-curr]
			}
		}
	}

	return dp[len(nums)][subsetSum]
}

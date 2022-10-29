package recursion

// https://leetcode.com/problems/permutations/

func Permute(nums []int) [][]int {
	var ans [][]int
	backtrackPermute(0, nums, &ans)
	return ans
}

func backtrackPermute(start int, nums []int, ans *[][]int) {

	if start == len(nums)-1 {
		*ans = append(*ans, append([]int{}, nums...))
	}

	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		backtrackPermute(start+1, nums, ans)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

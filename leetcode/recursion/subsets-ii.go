package recursion

import "sort"

// https://leetcode.com/problems/subsets-ii/

func SubsetsWithDup(nums []int) [][]int {

	ans := [][]int{{}}
	sort.Ints(nums)
	backtrackSubsetsWithDup(0, len(nums), nums, nil, &ans)
	return ans
}

func backtrackSubsetsWithDup(start, n int, nums, prev []int, ans *[][]int) {

	for i := start; i < n; i++ {
		if i > 0 && i > start && nums[i] == nums[i-1] {
			continue
		}
		prev = append(prev, nums[i])
		*ans = append(*ans, append([]int{}, prev...))
		backtrackSubsetsWithDup(i+1, n, nums, prev, ans)
		prev = prev[:len(prev)-1]
	}
}

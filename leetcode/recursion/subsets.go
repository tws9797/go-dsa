package recursion

// https://leetcode.com/problems/subsets/

func Subsets(nums []int) [][]int {

	ans := [][]int{{}}

	backtrackSubsets(0, len(nums), nums, nil, &ans)

	return ans
}

func backtrackSubsets(start, n int, nums, prev []int, ans *[][]int) {

	for i := start; i < n; i++ {
		prev = append(prev, nums[i])
		*ans = append(*ans, append([]int{}, prev...))
		backtrackSubsets(i+1, n, nums, prev, ans)
		prev = prev[:len(prev)-1]
	}
}

func SubsetsCascading(nums []int) [][]int {
	ans := [][]int{{}}

	for _, num := range nums {
		for _, existAns := range ans {
			temp := append([]int{}, existAns...)
			temp = append(temp, num)
			ans = append(ans, temp)
		}
	}

	return ans
}

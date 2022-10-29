package recursion

// https://leetcode.com/problems/combinations/

func Combine(n int, k int) [][]int {

	var ans [][]int
	backtrackCombinations(1, k, n, nil, &ans)

	return ans
}

/*
Input: n = 4, k = 2
Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
*/

func backtrackCombinations(start, k, n int, prev []int, result *[][]int) {
	if len(prev) == k {
		*result = append(*result, append([]int{}, prev...))
	}

	for i := start; i <= n; i++ {
		prev = append(prev, i)
		backtrackCombinations(i+1, k, n, prev, result)
		prev = prev[:len(prev)-1]
	}
}

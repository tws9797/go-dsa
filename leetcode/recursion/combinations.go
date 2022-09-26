package recursion

// https://leetcode.com/problems/combinations/

func Combine(n int, k int) [][]int {

	var ans [][]int
	backtrackCombinations(1, k, n, nil, &ans)

	return ans
}

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

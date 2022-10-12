package essentials

// https://leetcode.com/problems/subarray-sum-equals-k/

func SubarrayCumulativeSum(nums []int, k int) int {
	res := 0
	sum := make([]int, len(nums)+1)

	for i := 1; i <= len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	for start := 0; start < len(nums); start++ {
		for end := start + 1; end <= len(nums); end++ {
			if sum[end]-sum[start] == k {
				res++
			}
		}
	}

	return res
}

func SubarrayCumulativeSumWithoutSpace(nums []int, k int) int {

	res := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end < len(nums); end++ {
			sum += nums[end]
			if sum == k {
				res++
			}
		}
	}

	return res
}

func SubarraySum(nums []int, k int) int {

	tmp := make(map[int]int)
	tmp[0] = 1
	sum := 0
	res := 0
	// cumulative sum[i] - sum[j] = k if the sum between element i and j is k
	for _, v := range nums {
		sum += v
		if tmp[sum-k] > 0 {
			res += tmp[sum-k]
		}
		tmp[sum]++
	}

	return res
}

package recursion

// https://leetcode.com/problems/combination-sum/

func CombinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	ans := make([][]int, 0)
	doCombination(candidates, target, 0, nil, &ans)
	return ans
}

func doCombination(candidates []int, target int, prevSum int, prev []int, ans *[][]int) {
	if prevSum == target {
		*ans = append(*ans, append([]int{}, prev...))
		return
	}

	if prevSum > target {
		return
	}

	//2,3,5
	//8

	for i, v := range candidates {
		prev = append(prev, v)
		doCombination(candidates[i:], target, prevSum+v, prev, ans)
		prev = prev[:len(prev)-1]
	}
}

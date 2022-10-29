package recursion

func PermuteUnique(nums []int) [][]int {

	var ans [][]int
	BacktrackPermuteUnique(&ans, nums, 0)
	return ans
}

func BacktrackPermuteUnique(ans *[][]int, nums []int, start int) {

	if start == len(nums)-1 {
		*ans = append(*ans, append([]int{}, nums...))
		return
	}

	//1,1,2
	//[[1,1,2],[1,2,1],[1,1,2],[1,2,1],[2,1,1],[2,1,1]]
	/*
		[[1,1,2],
		 [1,2,1],
		 [2,1,1]]

	*/
	used := make(map[int]bool)

	for i := start; i < len(nums); i++ {
		if _, ok := used[nums[i]]; ok {
			continue
		}
		nums[start], nums[i] = nums[i], nums[start]
		BacktrackPermuteUnique(ans, nums, start+1)
		nums[start], nums[i] = nums[i], nums[start]
		used[nums[i]] = true
	}
}

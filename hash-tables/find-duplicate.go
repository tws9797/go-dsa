package hash_tables

func findDuplicate(nums []int) int {

	mapNums := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		if _, ok := mapNums[nums[i]]; ok {
			return nums[i]
		}

		mapNums[nums[i]] = true
	}

	return 0
}

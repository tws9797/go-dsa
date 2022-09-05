package arrays

// https://leetcode.com/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {

	// Create a hash table to store the new number encountered
	mapNums := map[int]bool{}

	for _, num := range nums {
		if _, ok := mapNums[num]; ok {
			return true
		}
		mapNums[num] = true
	}

	return false
}

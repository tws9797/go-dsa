package essentials

import "sort"

//https://leetcode.com/problems/valid-palindrome/

func ContainsDuplicate(nums []int) bool {

	// Create a hash table to store the new number encountered
	mapNums := map[int]struct{}{}

	for _, num := range nums {
		if _, ok := mapNums[num]; ok {
			return true
		}
		mapNums[num] = struct{}{}
	}

	return false
}

func ContainsDuplicate2(nums []int) bool {

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

package essentials

//https://leetcode.com/problems/two-sum/

func TwoSum(nums []int, target int) []int {

	// Store the complement and index pair in the hash table
	data := map[int]int{}

	for idx, num := range nums {

		// Check if the complement (target - num) exists in the hashmap
		if foundIdx, ok := data[target-num]; ok {

			// Iterating and inserting elements (complement and index pair) into the hash table
			return []int{foundIdx, idx}
		}

		data[num] = idx
	}

	return nil
}

func TwoSumBruteForce(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

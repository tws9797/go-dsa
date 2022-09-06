package arrays

// https://leetcode.com/problems/two-sum/
func TwoSum(nums []int, target int) []int {

	// Store the number to index in a hash table
	data := map[int]int{}

	for idx, num := range nums {

		// Store the num with their respective index in hash table
		// For each number, check if target - num exists in the hash table
		// If the hash table[target-num] is found, return the value (which is index) and the current index
		if foundNum, ok := data[target-num]; ok {
			return []int{foundNum, idx}
		}
		data[num] = idx
	}

	return []int{}
}

package essentials

// https://leetcode.com/problems/majority-element/
/*
	Input: nums = [2,2,1,1,1,2,2]
	Output: 2
*/

func MajorityElement(nums []int) int {

	m := map[int]int{}
	max := 0
	major := 0

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] > max {
			max = m[nums[i]]
			major = nums[i]
		}
	}

	return major
}

// Boyer-Moore Voting Algorithm
// Only works when one element is more than n/2

func MajorityElementSuffix(nums []int) int {

	// Maintain a count
	// Incremented whenever an instance of current candidate for majority element
	// Decremented when other element appear
	count := 0
	var candidate int

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

package essentials

// https://leetcode.com/problems/longest-consecutive-sequence/

func LongestConsecutive(nums []int) int {
	m := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}

	maxLength := 0

	for k, _ := range m {

		length := 1
		temp := k

		for _, ok := m[k+1]; ok; ok = m[k+1] {
			length++
			k = k + 1
			delete(m, k)
		}

		for _, ok := m[temp-1]; ok; ok = m[temp-1] {
			length++
			temp = temp - 1
			delete(m, temp)
		}

		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

func LongestConsecutive2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// longest represents the length of the longest consecutive elements sequence.
	longest := 1

	// set represents the set of numbers (removes duplicates).
	set := make(map[int]bool, len(nums))

	// add the numbers to the set.
	// O(n)
	for _, v := range nums {
		set[v] = true
	}

	// iterate through the set to find the longest consecutive sequence.
	for _, v := range nums {
		// if the previous number does not exist,
		// the current number represents the first number in a consecutive sequence.
		if !set[v-1] {

			// determine the length of the current consecutive sequence.
			current := 1
			v++
			for set[v] {
				current++
				v++
			}

			// set the longest consecutive sequence.
			if current > longest {
				longest = current
			}
		}
	}

	return longest
}

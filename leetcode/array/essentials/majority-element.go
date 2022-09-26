package essentials

// https://leetcode.com/problems/majority-element/

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

func MajorityElementSuffix(nums []int) int {

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

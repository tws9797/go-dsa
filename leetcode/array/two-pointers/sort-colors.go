package two_pointers

// https://leetcode.com/problems/sort-colors/

func SortColors(nums []int) {

	l := 0
	curr := 0
	r := len(nums) - 1

	for curr <= r {
		if nums[curr] == 0 {
			nums[curr], nums[l] = nums[l], nums[curr]
			l++
			curr++
		} else if nums[curr] == 2 {
			nums[curr], nums[r] = nums[r], nums[curr]
			r--
		} else {
			curr++
		}
	}
}

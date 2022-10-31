package essentials

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

func TwoSumInputArrayIsSorted(numbers []int, target int) []int {

	l, r := 0, len(numbers)-1
	sum := 0

	for l < r {

		sum = numbers[l] + numbers[r]

		if sum == target {
			return []int{l + 1, r + 1}
		}

		if sum < target {
			sum -= numbers[l]
			l++
		} else {
			sum -= numbers[r]
			r--
		}
	}

	return nil
}

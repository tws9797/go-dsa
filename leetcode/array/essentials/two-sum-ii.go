package essentials

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

func TwoSumInputArrayIsSorted(numbers []int, target int) []int {

	start, end := 0, len(numbers)-1

	for start < end {

		complement := target - numbers[start]

		for complement <= numbers[end] {

			if complement == numbers[end] {

				return []int{start + 1, end + 1}
			}

			end--
		}

		start++
	}

	return nil
}

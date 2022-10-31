package traversing_from_the_right

// https://leetcode.com/problems/number-of-visible-people-in-a-queue/

func CanSeePersonsCount(heights []int) []int {

	ans := make([]int, len(heights))

	// Keep a stack of decreasing values of heights
	var stack []int

	// Start from the end of the array (traversing from right)
	for curr := len(heights) - 1; curr >= 0; curr-- {

		currHeight := heights[curr]
		count := 0

		// If currHeight larger than top of the stack,
		// remove the top of the stack and replace it with currHeight
		for len(stack) > 0 && currHeight > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			count++
		}

		if len(stack) > 0 {
			count++
		}

		ans[curr] = count
		stack = append(stack, currHeight)
	}

	return ans
}

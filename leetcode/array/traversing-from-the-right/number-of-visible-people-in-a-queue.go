package traversing_from_the_right

// https://leetcode.com/problems/number-of-visible-people-in-a-queue/

func CanSeePersonsCount(heights []int) []int {

	lenHeights := len(heights)
	ans := make([]int, len(heights))

	// Keep a stack of decreasing values of heights
	var stack []int

	// Start from the end of the array (traversing from right)
	for curr := lenHeights - 1; curr >= 0; curr-- {

		currHeight := heights[curr]
		count := 0
		lenStack := len(stack)

		// If currHeight larger than top of the stack,
		// remove the top of the stack and replace it with currHeight
		for lenStack > 0 && currHeight > stack[lenStack-1] {
			stack = stack[:lenStack-1]
			lenStack--
			count++
		}

		if lenStack > 0 {
			count++
		}

		ans[curr] = count
		stack = append(stack, currHeight)
	}

	return stack
}

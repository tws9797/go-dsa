package leetcode

func CanSeePersonCount(heights []int) []int {

	lenHeights := len(heights)
	ans := make([]int, len(heights))

	// Keep a stack of decreasing values of heights
	var stack []int

	// Start from the end of the array
	for curr := lenHeights - 1; curr >= 0; curr-- {

		currHeight := heights[curr]
		count := 0
		lenStack := len(stack)

		// If currHeight larger than top of the stack,
		// remove the top of the stack and replace it with currHeight
		for lenStack > 0 && stack[lenStack-1] < currHeight {
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

	return ans
}

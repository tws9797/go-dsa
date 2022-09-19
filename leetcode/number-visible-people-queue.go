package leetcode

func CanSeePersonsCount(heights []int) []int {

	lenHeights := len(heights)
	ans := make([]int, lenHeights)
	ans[lenHeights-1] = 0
	ans[lenHeights-2] = 1
	highest := 0

	for i := lenHeights - 1; i >= 0; i-- {

		currHeight := heights[i]
		if currHeight > highest {
			highest = currHeight
			continue
		}

		nextHeight := 1
		for currHeight > heights[nextHeight+i] {
			nextHeight += ans[nextHeight+i]
		}
		ans[i] = nextHeight
	}

	return ans
}

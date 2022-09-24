package essentials

// https://leetcode.com/problems/container-with-most-water/

func MaxArea(heights []int) int {

	l, r := 0, len(heights)-1
	area := 0
	height := 0

	for l < r {

		width := r - l
		if heights[l] < heights[r] {
			height = heights[l]
			l++
		} else {
			height = heights[r]
			r--
		}

		if height*width > area {
			area = height * width
		}
	}

	return area
}

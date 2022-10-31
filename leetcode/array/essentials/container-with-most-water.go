package essentials

// https://leetcode.com/problems/container-with-most-water/
/*
	Input: height = [1,8,6,2,5,4,8,3,7]
	Output: 49
	Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
	In this case, the max area of water (blue section) the container can contain is 49.
*/

func MaxArea(heights []int) int {

	// Use two pointer to track the heights
	l, r := 0, len(heights)-1
	area := 0
	height := 0

	for l < r {

		width := r - l

		// To maximize the area, the height of two side must as high as possible
		// Move the shorter line, to check for the higher line, to get more area
		if heights[l] < heights[r] {
			height = heights[l]
			l++
		} else {
			height = heights[r]
			r--
		}

		// Maintain a maxarea to store the maximum area obtained till now
		if height*width > area {
			area = height * width
		}
	}

	return area
}

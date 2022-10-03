package matrix

// https://leetcode.com/problems/spiral-matrix/

func SpiralOrder(matrix [][]int) []int {

	//Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
	//Output: [1,2,3,6,9,8,7,4,5]

	R := len(matrix)
	C := len(matrix[0])
	var ans []int
	required := R * C
	rPtr, cPtr := 0, 0
	var right, left, up, down bool
	right = true

	for len(ans) != required {

		if right {

			for i := 0; i < C; i++ {
				ans = append(ans, matrix[rPtr][cPtr])
				cPtr++
			}

			R--
			rPtr++
			cPtr--
			right = false
			down = true
		} else if down {

			for i := 0; i < R; i++ {
				ans = append(ans, matrix[rPtr][cPtr])
				rPtr++
			}

			C--
			rPtr--
			cPtr--
			down = false
			left = true
		} else if left {

			for i := 0; i < C; i++ {
				ans = append(ans, matrix[rPtr][cPtr])
				cPtr--
			}

			R--
			rPtr--
			cPtr++
			left = false
			up = true
		} else if up {

			for i := 0; i < R; i++ {
				ans = append(ans, matrix[rPtr][cPtr])
				rPtr--
			}

			C--
			rPtr++
			cPtr++
			up = false
			right = true
		}
	}

	return ans
}

func SpiralOrder2(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}
	rowBegin, rowEnd, colBegin, colEnd := 0, len(matrix)-1, 0, len(matrix[0])-1

	for rowBegin <= rowEnd && colBegin <= colEnd {
		for i := colBegin; i <= colEnd; i++ {
			res = append(res, matrix[rowBegin][i])
		}
		rowBegin++

		for i := rowBegin; i <= rowEnd; i++ {
			res = append(res, matrix[i][colEnd])
		}
		colEnd--

		if rowBegin <= rowEnd {
			for i := colEnd; i >= colBegin; i-- {
				res = append(res, matrix[rowEnd][i])
			}
		}
		rowEnd--

		if colBegin <= colEnd {
			for i := rowEnd; i >= rowBegin; i-- {
				res = append(res, matrix[i][colBegin])
			}
		}
		colBegin++
	}

	return res
}

func SpiralOrder3(matrix [][]int) []int {

	if len(matrix) == 0 {
		return []int{}
	}
	var res []int
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	size := len(matrix) * len(matrix[0])

	for len(res) < size {
		for i := left; i <= right && len(res) < size; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		for i := top; i <= bottom && len(res) < size; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		for i := right; i >= left && len(res) < size; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--
		for i := bottom; i >= top && len(res) < size; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}

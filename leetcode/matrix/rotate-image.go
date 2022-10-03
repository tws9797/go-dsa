package matrix

// https://leetcode.com/problems/rotate-image/

func Rotate(matrix [][]int) {

	n := len(matrix) - 1
	for x := 0; x < len(matrix)/2; x++ {
		for y := x; y < n-x; y++ {
			// save top left
			t := matrix[x][y]
			// set top left to bottom left
			matrix[x][y] = matrix[n-y][x]
			// set bottom left to bottom right
			matrix[n-y][x] = matrix[n-x][n-y]
			// set bottom right to top right
			matrix[n-x][n-y] = matrix[y][n-x]
			// set top right to top left
			matrix[y][n-x] = t
		}
	}
}

func RotateTransposeReverse(matrix [][]int) {
	// diagonal symmetry change
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// column symmetry change
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[j][i], matrix[j][len(matrix)-1-i] = matrix[j][len(matrix)-1-i], matrix[j][i]
		}
	}
}

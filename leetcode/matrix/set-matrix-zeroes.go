package matrix

// https://leetcode.com/problems/set-matrix-zeroes/

func SetZeroes(matrix [][]int) {
	R := len(matrix)
	C := len(matrix[0])
	rows := map[int]bool{}
	cols := map[int]bool{}

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func SetZeroesSpaceEfficient(matrix [][]int) {

	var isCol bool
	R := len(matrix)
	C := len(matrix[0])

	for i := 0; i < R; i++ {

		// Since first cell for both first row and first column is the same i.e. matrix[0][0]
		// We can use an additional variable for either the first row/column.
		// For this solution we are using an additional variable for the first column
		// and using matrix[0][0] for the first row.

		if matrix[i][0] == 0 {
			isCol = true
		}

		// Use the first cell of every row and column as a flag. The flag would determine whether a row or column has been set to zero.
		for j := 1; j < C; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// Iterate over the array once again and using the first row and first column, update the elements.
	for i := 1; i < R; i++ {
		for j := 1; j < C; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// See if the first row needs to be set to zero as well
	if matrix[0][0] == 0 {
		for j := 0; j < C; j++ {
			matrix[0][j] = 0
		}
	}

	// See if the first column needs to be set to zero as well
	if isCol {
		for i := 0; i < R; i++ {
			matrix[i][0] = 0
		}
	}
}

package math

// https://leetcode.com/problems/sqrtx/

func MySqrt(x int) int {
	if x < 2 {
		return x
	}

	// For x >= 2, the square root is always smaller than or equal to x/2 and larger than 0
	left, right := 2, x/2
	pivot := 0

	for left <= right {

		pivot = left + (right-left)/2
		num := pivot * pivot

		if num == x {
			return pivot
		} else if num > x {
			right = pivot - 1
		} else {
			left = pivot + 1
		}

	}

	return right
}

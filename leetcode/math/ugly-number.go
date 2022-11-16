package math

// https://leetcode.com/problems/ugly-number/
/*
	Input: n = 6
	Output: true
	Explanation: 6 = 2 × 3

	Input: n = 14
	Output: false
	Explanation: 14 is not ugly since it includes the prime factor 7.
*/

func IsUgly(n int) bool {
	// An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.
	// A non-positive integer cannot be ugly
	if n <= 0 {
		return false
	}

	// Factorize by dividing with permitted factors.
	for _, factor := range []int{2, 3, 5} {
		n = keepDividingWhenDivisible(n, factor)
	}

	// Check if the integer is reduced to 1 or not.
	return n == 1
}

// Keep dividing dividend by divisor when division is possible.
func keepDividingWhenDivisible(dividend, divisor int) int {

	for dividend%divisor == 0 {
		dividend /= divisor
	}

	return dividend
}

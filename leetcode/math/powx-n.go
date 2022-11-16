package math

// https://leetcode.com/problems/powx-n/

func MyPow(x float64, n int) float64 {

	/*
		if brute force, 10 times executed
		2^10 = 2*2*2*2*2*2*2*2*2*2

		if fast power, 5 times executed
		2^10 = 2^5 * 2^5
		2^5 = 2^2 * 2^2 * 2
		2^2 = 2 * 2
	*/
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x, n = 1/x, -n
	}
	if n%2 == 0 {
		return MyPow(x*x, n/2)
	}
	return MyPow(x*x, n/2) * x
}

func MyPowIterative(x float64, n int) float64 {

	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n < 0 {
		x, n = 1/x, -n
	}

	ans := 1.0

	for i := n; i > 0; i /= 2 {
		if i%2 == 1 {
			ans *= x
		}

		x *= x
	}

	return ans
}

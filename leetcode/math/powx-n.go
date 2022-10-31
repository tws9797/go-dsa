package math

func MyPow(x float64, n int) float64 {
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

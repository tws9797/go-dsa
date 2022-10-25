package traversing_from_the_right

// https://leetcode.com/problems/daily-temperatures/

func DailyTemperatures(temperatures []int) []int {

	lenTemp := len(temperatures)
	ans := make([]int, lenTemp)
	ans[lenTemp-1] = 0

	for i := lenTemp - 2; i >= 0; i-- {

		if temperatures[i] < temperatures[i+1] {
			ans[i] = 1
		} else {
			nextPointer := i + 1
			for ans[nextPointer] != 0 {
				nextPointer += ans[nextPointer]
				nextTemperature := temperatures[nextPointer]
				if nextTemperature > temperatures[i] {
					ans[i] = nextPointer - i
					break
				}
			}
		}
	}

	return ans
}

func DailyTemperaturesMonotonicStack(temperatures []int) []int {

	n := len(temperatures)
	ans := make([]int, n)

	// Monotonic stack is simply a stack where the elements are always in sorted order
	// Monotonic decreasing means that the stack will always be sorted in decreasing order
	var stack []int

	for currDay := 0; currDay < n; currDay++ {
		currentTemp := temperatures[currDay]

		// Compare current temperature to the top of the stack
		for len(stack) != 0 && currentTemp > temperatures[stack[len(stack)-1]] {

			// If current temperature is greater than the top of the stock
			// Pop the stack until it reach the warmest temperature
			prevDay := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevDay] = currDay - prevDay
		}

		stack = append(stack, currDay)
	}

	return ans
}

func DailyTemperaturesArray(temperatures []int) []int {

	n := len(temperatures)
	ans := make([]int, n)
	hottest := 0

	for currDay := n - 1; currDay >= 0; currDay-- {
		currTemp := temperatures[currDay]
		if currTemp >= hottest {
			hottest = currTemp
			continue
		}

		//73,74,75,71,69,72,76,73
		days := 1
		for currTemp >= temperatures[currDay+days] {
			days += ans[currDay+days]
		}
		ans[currDay] = days
	}

	return ans
}

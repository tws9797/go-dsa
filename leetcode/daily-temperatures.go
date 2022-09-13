package leetcode

func DailyTemperatures(temperatures []int) []int {

	res := make([]int, len(temperatures))

	//+1,+1,+1,-4,-2,-3,+3,+4,-3
	//73,74,75,71,69,72,76,73
	//1,1,4,2,1,1,0,0
	res[len(temperatures)-1] = 0
	for i := len(temperatures) - 2; i >= 0; i-- {

	}

	return res
}

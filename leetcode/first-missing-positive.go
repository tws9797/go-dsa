package leetcode

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func FirstMissingPositive(nums []int) int {

	lenNums := len(nums)
	contains := 0

	for i := 0; i < lenNums; i++ {
		if nums[i] == 1 {
			contains++
			break
		}
	}

	if contains == 0 {
		return 1
	}

	for i := 0; i < lenNums; i++ {
		if nums[i] <= 0 || nums[i] > lenNums {
			nums[i] = 1
		}
	}

	for i := 0; i < lenNums; i++ {

		a := Abs(nums[i])

		if a == lenNums {
			nums[0] = -Abs(nums[0])
		} else {
			nums[a] = -Abs(nums[a])
		}
	}

	for i := 1; i < lenNums; i++ {
		if nums[i] > 0 {
			return i
		}
	}

	if nums[0] > 0 {
		return lenNums
	}

	return lenNums + 1
}

func FirstMissingPositiveSort(nums []int) int {

	i := 0
	lenNums := len(nums)

	for i < lenNums {
		j := nums[i] - 1
		if 0 <= j && j < lenNums && nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i += 1
		}
	}

	for i, val := range nums {
		if val != i+1 {
			return i + 1
		}
	}

	return lenNums + 1
}

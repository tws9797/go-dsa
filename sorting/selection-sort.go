package sorting

import "math"

func SelectionSort(arr []int) {

	minIdx := 0

	for i := 0; i < len(arr); i++ {

		min := math.MaxInt

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min {
				minIdx = j
				min = arr[j]
			}
		}

		arr[i], arr[minIdx] = min, arr[i]
	}
}

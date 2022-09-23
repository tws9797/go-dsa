package sorting_the_array

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/non-overlapping-intervals/

func EraseOverlapIntervals(intervals [][]int) int {

	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	fmt.Println(intervals)

	count := 1
	end := intervals[0][1]
	//[[1 2] [2 3] [1 5] [3 6]]

	//[[1,2] [1,3] [2,3] [3,4]]

	for i := 1; i < len(intervals); i++ {

		if intervals[i][0] >= end {
			end = intervals[i][1]
			count++
		}

	}

	return len(intervals) - count
}

package leetcode

import "sort"

func MergeIntervals(intervals [][]int) [][]int {
	var mergeIntervals [][]int
	start, end := 0, 0
	lenIntervals := len(intervals)
	i := 0

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	mergeIntervals = append(mergeIntervals, []int{intervals[0][0], intervals[0][1]})

	for i < lenIntervals-1 {
		start = mergeIntervals[len(mergeIntervals)-1][0]
		end = mergeIntervals[len(mergeIntervals)-1][1]

		if end < intervals[i+1][0] {
			mergeIntervals = append(mergeIntervals, []int{intervals[i+1][0], intervals[i+1][1]})
		} else {
			if start > intervals[i+1][0] {
				start = intervals[i+1][0]
			}

			if end < intervals[i+1][1] {
				end = intervals[i+1][1]
			}

			mergeIntervals[len(mergeIntervals)-1][0] = start
			mergeIntervals[len(mergeIntervals)-1][1] = end
		}

		i++
	}

	return mergeIntervals
}

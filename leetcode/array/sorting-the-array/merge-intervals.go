package sorting_the_array

import "sort"

// https://leetcode.com/problems/merge-intervals/
/*
	Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
	Output: [[1,6],[8,10],[15,18]]
	Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
*/

func MergeIntervals(intervals [][]int) [][]int {

	var mergeIntervals [][]int

	// Sort the starting value of each array
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Add the first array to the mergeIntervals
	mergeIntervals = append(mergeIntervals, intervals[0])
	for i := 1; i < len(intervals); i++ {

		lenMergeIntervals := len(mergeIntervals)

		if intervals[i][0] <= mergeIntervals[lenMergeIntervals-1][1] {
			if intervals[i][1] > mergeIntervals[lenMergeIntervals-1][1] {
				mergeIntervals[lenMergeIntervals-1][1] = intervals[i][1]
			}
		} else {
			mergeIntervals = append(mergeIntervals, intervals[i])
		}
	}

	return mergeIntervals
}

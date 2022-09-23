package sorting_the_array

import "sort"

// https://leetcode.com/problems/merge-intervals/

func MergeIntervals(intervals [][]int) [][]int {

	var mergeIntervals [][]int
	end := 0
	lenIntervals := len(intervals)
	i := 0

	// Sort the starting value of each array
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Add the first array to the mergeIntervals
	mergeIntervals = append(mergeIntervals, []int{intervals[0][0], intervals[0][1]})
	//[[1,3],[2,6],[8,10],[15,18]]
	for i < lenIntervals-1 {
		end = mergeIntervals[len(mergeIntervals)-1][1]

		if end < intervals[i+1][0] {
			mergeIntervals = append(mergeIntervals, []int{intervals[i+1][0], intervals[i+1][1]})
		} else {
			if end < intervals[i+1][1] {
				end = intervals[i+1][1]
			}

			mergeIntervals[len(mergeIntervals)-1][1] = end
		}

		i++
	}

	return mergeIntervals
}

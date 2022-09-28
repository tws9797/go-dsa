package intervals

import (
	"sort"
)

//Input: intervals = [[1,2],[2,3],[2,4],[1,3]]
//Sort by starting points = [[1,2],[1,3],[2,3],[2,4]]
//Sort by ending points = [[1,2],[2,3],[1,3],[2,4]]
//Output: 1

// https://leetcode.com/problems/non-overlapping-intervals/

func EraseOverlapIntervalsByStartPoint(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Sort so that the starting point of each interval is in order
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	count := 0
	prev := 0

	/*
			--
			---
			 --
		     ---
	*/

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[prev][1] {
			if intervals[i][1] < intervals[prev][1] {
				prev = i
			}
			count++
		} else {
			prev = i
		}
	}

	return count
}

func EraseOverlapIntervalsEndPoint(intervals [][]int) int {

	if len(intervals) == 0 {
		return 0
	}

	// Sort so that the ending point of each interval is in order
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 1
	end := intervals[0][1]

	/*
			--
		     --
		    ---
		     ---
	*/

	for i := 1; i < len(intervals); i++ {

		if intervals[i][0] >= end {
			end = intervals[i][1]
			count++
		}
	}

	return len(intervals) - count
}

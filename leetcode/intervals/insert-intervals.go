package intervals

import (
	"sort"
)

// https://leetcode.com/problems/insert-interval/

func Insert(intervals [][]int, newInterval []int) [][]int {

	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	mergeIntervals := [][]int{intervals[0]}
	end := 0

	for i := 0; i < len(intervals)-1; i++ {
		end = mergeIntervals[len(mergeIntervals)-1][1]

		if end < intervals[i+1][0] {
			mergeIntervals = append(mergeIntervals, intervals[i+1])
		} else {
			if end < intervals[i+1][1] {
				end = intervals[i+1][1]
			}

			mergeIntervals[len(mergeIntervals)-1][1] = end
		}
	}

	return mergeIntervals
}

/*
Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
*/

func Insert2(intervals [][]int, newInterval []int) [][]int {
	var l [][]int
	var r [][]int
	a, b := newInterval[0], newInterval[1]

	for _, interval := range intervals {
		if interval[1] < a {
			l = append(l, interval)
		} else if interval[0] > b {
			r = append(r, interval)
		} else {
			a = min(a, interval[0])
			b = max(b, interval[1])
		}
	}

	return append(append(l, []int{a, b}), r...)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

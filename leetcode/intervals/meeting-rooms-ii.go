package intervals

import "sort"

// https://leetcode.com/problems/meeting-rooms-ii/

func minMeetingRooms(intervals [][]int) int {
	/*
		Input: intervals = [[0,30],[5,16],[15,20][30,40],[35,36]]
		Output: 2
	*/
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	count := 0

	for i := 0; i < len(intervals); i++ {

		if intervals[i][1] > intervals[i+1][0] {
			count++
		}
	}

	return count
}

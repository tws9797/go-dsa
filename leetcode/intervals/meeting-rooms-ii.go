package intervals

import "sort"

// https://leetcode.com/problems/meeting-rooms-ii/

func minMeetingRooms(intervals [][]int) int {

	lenIntervals := len(intervals)

	starts := make([]int, lenIntervals)
	ends := make([]int, lenIntervals)

	for i := 0; i < lenIntervals; i++ {
		starts[i] = intervals[i][0]
		ends[i] = intervals[i][1]
	}

	sort.Ints(starts)
	sort.Ints(ends)

	startPtr := 0
	endPtr := 0
	count := 0

	for startPtr < lenIntervals {
		if starts[startPtr] < ends[endPtr] {
			count++
		} else {
			endPtr++
		}

		startPtr++
	}

	return count
}

// https://leetcode.com/problems/meeting-rooms/description/
package meetingrooms

import "sort"

type Interval struct {
	Start, End int
}

func CanAttendMeetings(intervals []*Interval) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start < intervals[i-1].End {
			return false
		}
	}
	return true
}

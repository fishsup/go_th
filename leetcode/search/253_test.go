package search

import (
	"sort"
	"testing"
)

func TestMeestingRooms(t *testing.T){
	intervals := [][2]int{}
	intervals = append(intervals, [2]int{2,3})
	intervals = append(intervals, [2]int{1,3})
	t.Log(minMeetingRooms(intervals))
}

func minMeetingRooms(intervals [][2]int) int {
	if len(intervals)==0 {
		return 0
	}
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	return -1;
}
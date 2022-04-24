package recursion

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	rs := [][]int{}
	previous := []int{}
	for _, interval := range intervals {
		if len(previous) == 0 || interval[0] > previous[1] {
			if len(previous) != 0 {
				rs = append(rs, previous)
			}
			previous = interval
		} else {
			tail := previous[1]
			if interval[1] > tail {
				tail = interval[1]
			}
			previous[1] = tail
		}
	}
	if len(previous) != 0 {
		rs = append(rs, previous)
	}
	return rs
}

// @lc code=end

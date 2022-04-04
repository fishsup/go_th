package search

import "testing"

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

func TestA(t *testing.T) {
	t.Log(searchRange([]int{1}, 1))
}

// @lc code=start
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || target < nums[0] || target > nums[len(nums)-1] {
		return []int{-1, -1}
	}
	return []int{searchFirstEle(nums, target, 0, len(nums)-1), searchLastEle(nums, target, 0, len(nums)-1)}
}

func searchFirstEle(nums []int, target, low, hign int) int {
	if low > hign {
		return -1
	}
	mid := low + (hign-low)/2
	if nums[mid] == target && (mid == 0 || nums[mid-1] < target) {
		return mid
	}
	if nums[mid] < target {
		return searchFirstEle(nums, target, mid+1, hign)
	} else {
		return searchFirstEle(nums, target, low, mid-1)
	}
}

func searchLastEle(nums []int, target, low, hign int) int {
	if low > hign {
		return -1
	}
	mid := low + (hign-low)/2
	if nums[mid] == target && (mid == len(nums)-1 || nums[mid+1] > target) {
		return mid
	}
	if nums[mid] > target {
		return searchLastEle(nums, target, low, mid-1)
	} else {
		return searchLastEle(nums, target, mid+1, hign)
	}
}
// @lc code=end
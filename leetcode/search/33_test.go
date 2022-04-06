package search

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 *
 */

/*
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
*/

// @lc code=start
func search_33(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return binarySearch33(nums, target, 0, len(nums)-1)
}

func binarySearch33(nums []int, target, low, hign int) int {
	if low > hign {
		return -1
	}
	middle := low + (hign-low)/2
	if nums[middle] == target {
		return middle
	}

	if nums[middle] >= nums[low] {
		//左侧有序
		if nums[middle] > target && nums[low] <= target {
			return binarySearch33(nums, target, low, middle-1)
		}
		return binarySearch33(nums, target, middle+1, hign)
	} else {
		//右侧有序
		if nums[middle] < target && nums[hign] >= target {
			return binarySearch33(nums, target, middle+1, hign)
		}
		return binarySearch33(nums, target, low, middle-1)
	}
}

// @lc code=end

package dp

import (
	"fmt"
	"testing"
)

/**
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。


示例 1：

输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
示例 2：

输入：nums = [0,1,0,3,2,3]
输出：4
示例 3：

输入：nums = [7,7,7,7,7,7,7]
输出：1


提示：

1 <= nums.length <= 2500
-104 <= nums[i] <= 104


进阶：

你能将算法的时间复杂度降低到 O(n log(n)) 吗?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// f(n) = max(f(i))+1 (1<=i<n-1 && nums[i-1]<nums[n-1])

var rs int
var cache map[int]int

func lengthOfLIS(nums []int) int {
	rs = 0
	if len(nums) != 0 {
		cache = map[int]int{}
		dp1(nums, len(nums)-1)
	}
	return rs
}

func dp1(nums []int, index int) int {
	if index < 0 {
		return 0
	}
	maxEnding := 1
	for i := 0; i < index; i++ {
		var t int
		if cachers, exist := cache[i]; exist {
			t = cachers
		} else {
			t = dp1(nums, i)
		}
		if nums[i] < nums[index] && t+1 > maxEnding {
			maxEnding = t + 1
		}
	}
	if maxEnding > rs {
		rs = maxEnding
	}
	cache[index] = maxEnding
	return maxEnding
}

func TestLe(t *testing.T) {
	// var b = lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10})
	var b = lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})
	fmt.Print(b)
}

package deque

import "testing"

//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//。
//
// 返回 滑动窗口中的最大值 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// 示例 2：
//
//
//输入：nums = [1], k = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
// 1 <= k <= nums.length
//
// Related Topics 队列 数组 滑动窗口 单调队列 堆（优先队列） 👍 1391 👎 0
func TestName(t *testing.T) {
	maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3)
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {
	//定义滑动窗口 存储窗口内数组下标
	window := []int{}
	result := []int{}
	push := func(i int) {
		//队尾值小于当前值则出队
		for len(window) > 0 && nums[window[len(window)-1]] < nums[i] {
			window = window[:len(window)-1]
		}
		window = append(window, i)
	}
	for i, _ := range nums {
		if i < k {
			//窗口初始化情况
			push(i)
			//达到窗口大小后设置第一个最大值
			if i == k-1 {
				result = append(result, nums[window[0]])
			}
		} else {
			//头部出队
			if i-window[0] > k-1 {
				window = window[1:]
			}
			//右侧值队尾入队
			push(i)
			result = append(result, nums[window[0]])
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

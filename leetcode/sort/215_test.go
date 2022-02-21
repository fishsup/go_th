package sort

import (
	"math/rand"
	"testing"
	"time"
)

//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
//
//
// 示例 1:
//
//
//输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
//
//
// 示例 2:
//
//
//输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4
//
//
//
// 提示：
//
//
// 1 <= k <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 1485 👎 0

func TestFindKthLargest(t *testing.T) {
	//arr := []int{2, 1, 7, 9, 5, 8}
	//arr := []int{3, 2, 1, 5, 6, 4}
	arr := []int{3, 1, 2, 4}
	t.Log(findKthLargest(arr, 2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return qsort215(nums, 0, len(nums)-1, k)
}

func qsort215(nums []int, start, end, k int) int {
	mid := rand.Int()%(end-start+1) + start
	nums[mid], nums[start] = nums[start], nums[mid]
	m := start
	d := nums[start]
	for i := start + 1; i <= end; i++ {
		if nums[i] > d {
			nums[m+1], nums[i] = nums[i], nums[m+1]
			m++
		}
	}
	nums[m], nums[start] = nums[start], nums[m]
	if m-start+1 > k {
		return qsort215(nums, start, m-1, k)
	} else if m-start+1 == k {
		return nums[m]
	} else {
		return qsort215(nums, m+1, end, k-m+start-1)
	}
}

//leetcode submit region end(Prohibit modification and deletion)

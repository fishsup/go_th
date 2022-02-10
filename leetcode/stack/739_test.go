package stack

import "testing"

//请根据每日 气温 列表 temperatures ，请计算在每一天需要等几天才会有更高的温度。如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
// 示例 1:
//
//
//输入: temperatures = [73,74,75,71,69,72,76,73]
//输出: [1,1,4,2,1,1,0,0]
//
//
// 示例 2:
//
//
//输入: temperatures = [30,40,50,60]
//输出: [1,1,1,0]
//
//
// 示例 3:
//
//
//输入: temperatures = [30,60,90]
//输出: [1,1,0]
//
//
//
// 提示：
//
//
// 1 <= temperatures.length <= 10⁵
// 30 <= temperatures[i] <= 100
//
// Related Topics 栈 数组 单调栈 👍 1011 👎 0
func TestName(t *testing.T) {
	t.Log(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func dailyTemperatures(temperatures []int) []int {
	results := make([]int, len(temperatures))
	stack := []int{}
	for i, val := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			//取栈顶位置对应温度值比较
			headVal := temperatures[stack[len(stack)-1]]
			for len(stack) >= 1 && val > headVal {
				head := stack[len(stack)-1]
				results[head] = i - head
				stack = stack[:len(stack)-1]
				if len(stack) >= 1 {
					headVal = temperatures[stack[len(stack)-1]]
				}
			}
			stack = append(stack, i)
		}
	}
	for _, p := range stack {
		results[p] = 0
	}
	return results
}

//leetcode submit region end(Prohibit modification and deletion)

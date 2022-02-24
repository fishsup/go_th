package backtracking

import "testing"

/*
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，
并以列表形式返回。你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

对于给定的输入，保证和为 target 的不同组合数少于 150 个。



示例 1：

输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。
示例 2：

输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]
示例 3：

输入: candidates = [2], target = 1
输出: []


提示：

1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都 互不相同
1 <= target <= 500


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

func TestCom(t *testing.T) {
	t.Log(combinationSum([]int{2, 3, 6, 7}, 7))
}

var rs [][]int

func combinationSum(candidates []int, target int) [][]int {
	rs = [][]int{}
	backtracking(candidates, target, 0, []int{})
	return rs
}

func backtracking(candidates []int, target, start int, solution []int) {
	if target < 0 {
		return
	}
	if target == 0 {
		rs = append(rs, append([]int(nil), solution...))
		return
	}
	for i := start; i < len(candidates); i++ {
		solution = append(solution, candidates[i])
		backtracking(candidates, target-candidates[i], i, solution)
		solution = solution[:len(solution)-1]
	}
}

func combinationSum1(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(target int, idx int)
	dfs = func(target int, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		dfs(target, idx+1)
		if target-candidates[idx]>=0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return ans
}

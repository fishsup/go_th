package backtracking

import (
	"sort"
	"testing"
)

/*
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用 一次 。

注意：解集不能包含重复的组合。



示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
输出:
[
[1,2,2],
[5]
]


提示:

1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30
*/
func TestA(t *testing.T) {
	t.Log(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	numtable := [][2]int{}
	for _, num := range candidates {
		if len(numtable) == 0 || numtable[len(numtable)-1][0] != num {
			numtable = append(numtable, [2]int{num, 1})
		} else {
			numtable[len(numtable)-1][1]++
		}
	}

	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if target == 0 {
			ans = append(ans, append([]int{}, comb...))
			return
		}
		//idx上没有数字或者target小于idx上的数字直接返回
		if idx == len(numtable) || target < numtable[idx][0] {
			return
		}
		dfs(target, idx+1)
		//计算当前数字最大可组合次数
		times := target / numtable[idx][0]
		if times > numtable[idx][1] {
			times = numtable[idx][1]
		}
		for i := 1; i <= times; i++ {
			comb = append(comb, numtable[idx][0])
			dfs(target-i*numtable[idx][0], idx+1)
		}
		comb = comb[:len(comb)-times]
	}
	dfs(target, 0)
	return ans
}


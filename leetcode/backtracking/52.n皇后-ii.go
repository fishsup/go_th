package backtracking

/*
n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。



示例 1：


输入：n = 4
输出：2
解释：如上图所示，4 皇后问题存在两个不同的解法。
示例 2：

输入：n = 1
输出：1


提示：

1 <= n <= 9


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-queens-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func totalNQueens(n int) int {
	rs := 0
	//列上是否有皇后
	columns := make([]bool, n)
	diagonal1 := make([]bool, 2*n-1)
	diagonal2 := make([]bool, 2*n-1)

	var backtrackingQ func(int)
	backtrackingQ = func(row int) {
		if row == n {
			rs++
			//已有列都已经分配了皇后
			return
		}
		for col, hasQueue := range columns {
			if hasQueue || diagonal1[col+row] || diagonal2[col-row+n-1] {
				//列或对角线上已经分配了皇后 直接跳过
				continue
			}
			diagonal1[col+row] = true
			diagonal2[col-row+n-1] = true
			columns[col] = true
			backtrackingQ(row + 1)
			diagonal1[col+row] = false
			diagonal2[col-row+n-1] = false
			columns[col] = false
		}
	}
	backtrackingQ(0)
	return rs
}

package backtracking

import (
	"strings"
	"testing"
)

/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */
func TestN(t *testing.T) {
	t.Log(solveNQueens(4))
}

// @lc code=start
func solveNQueens(n int) [][]string {
	rs := [][]string{}

	solution := make([][]string, n)
	for i := 0; i < n; i++ {
		solution[i] = make([]string, n)
		for j := 0; j < n; j++ {
			solution[i][j] = "."
		}
	}

	columns := make([]bool, n)
	diagonal1 := make([]bool, 2*n-1)
	diagonal2 := make([]bool, 2*n-1)

	var backtrackingQ1 func(int)
	backtrackingQ1 = func(row int) {
		if row == n {
			s := make([]string, n)
			for i := 0; i < n; i++ {
				s[i] = strings.Join(solution[i], "")
			}
			//todo rs
			rs = append(rs, s)
		}
		for col, hasQueue := range columns {
			if hasQueue || diagonal1[col+row] || diagonal2[col-row+n-1] {
				continue
			}
			columns[col] = true
			diagonal1[col+row] = true
			diagonal2[col-row+n-1] = true
			solution[row][col] = "Q"
			backtrackingQ1(row + 1)
			solution[row][col] = "."
			columns[col] = false
			diagonal1[col+row] = false
			diagonal2[col-row+n-1] = false
		}
	}

	backtrackingQ1(0)
	return rs
}

// @lc code=end

package trie //给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words， 返回所有二维网格上的单词 。
import (
	"testing"
)

//
// 单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使
//用。
//
//
//
// 示例 1：
//
//
//输入：board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f",
//"l","v"]], words = ["oath","pea","eat","rain"]
//输出：["eat","oath"]
//
//
// 示例 2：
//
//
//输入：board = [["a","b"],["c","d"]], words = ["abcb"]
//输出：[]
//
//
//
//
// 提示：
//
//
// m == board.length
// n == board[i].length
// 1 <= m, n <= 12
// board[i][j] 是一个小写英文字母
// 1 <= words.length <= 3 * 10⁴
// 1 <= words[i].length <= 10
// words[i] 由小写英文字母组成
// words 中的所有字符串互不相同
//
// Related Topics 字典树 数组 字符串 回溯 矩阵 👍 619 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
var rs map[string]string

func findWords(board [][]byte, words []string) []string {
	//构建trie
	trieMap := map[interface{}]interface{}{}
	for _, word := range words {
		cMap := trieMap
		charArr := []byte(word)
		for i, chara := range charArr {
			if a, ok := cMap[chara]; ok {
				cMap = a.(map[interface{}]interface{})
			} else {
				cMap[chara] = map[interface{}]interface{}{}
				cMap = cMap[chara].(map[interface{}]interface{})
			}
			if i == len(charArr)-1 {
				//结尾符号
				cMap['.'] = nil
			}
		}
	}
	//遍历二维数组
	rs = map[string]string{}
	widge := len(board)
	col := len(board[0])
	for i := 0; i < widge; i++ {
		for j := 0; j < col; j++ {
			//开始节点为i,j
			str := []byte{}
			dfs(i, j, board, trieMap, str)
		}
	}
	keys := make([]string, len(rs))
	j := 0
	for key := range rs {
		keys[j] = key
		j++
	}
	return keys
}

var directionArr = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func dfs(i, j int, board [][]byte, treeMap map[interface{}]interface{}, str []byte) {
	if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) && board[i][j] != '#' {
		charrrr := board[i][j]
		if nMap, ok := treeMap[charrrr]; ok {
			str = append(str, charrrr)
			//如果有结果标识符代表匹配到了一个单词
			if _, ok1 := nMap.(map[interface{}]interface{})['.']; ok1 {
				rs[string(str)] = ""
			}
			for k := 0; k < len(directionArr); k++ {
				board[i][j] = '#'
				dfs(directionArr[k][0]+i, directionArr[k][1]+j, board, nMap.(map[interface{}]interface{}), str)
				board[i][j] = charrrr
			}
		}
	}
}

func TestName1(t *testing.T) {
	//board := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	//words := []string{"oath", "pea", "eat", "rain"}

	board := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	words := []string{"oath", "pea", "eat", "rain", "hklf", "hf"}

	//测试结果:["oath","eat","hf"]
	//期望结果:["oath","eat","hklf","hf"]
	findWords(board, words)
}

func TestName(t *testing.T) {
	words := []string{"123", "1234", "12345", "1233", "213"}
	trieMap := map[interface{}]interface{}{}
	for _, word := range words {
		cMap := trieMap
		charArr := []byte(word)
		for i, chara := range charArr {
			if a, ok := cMap[chara]; ok {
				cMap = a.(map[interface{}]interface{})
			} else {
				cMap[chara] = map[interface{}]interface{}{}
				cMap = cMap[chara].(map[interface{}]interface{})
			}
			if i == len(charArr)-1 {
				//结尾符号
				cMap['.'] = nil
			}
		}
	}
	t.Log(trieMap)
}

//leetcode submit region end(Prohibit modification and deletion)

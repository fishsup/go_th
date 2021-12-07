package leetcode242

import "testing"

/* 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false


提示:

1 <= s.length, t.length <= 5 * 104
s 和 t 仅包含小写字母

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-anagram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

func TestLeetCode(t *testing.T) {
	t.Log(isAnagram("tese", "testsasdf"))
	t.Log(isAnagram("test", "test"))
}

func isAnagram(s string, t string) bool {
	var carr [26]int
	for _, chart := range s {
		carr[chart-'a']++
	}
	for _, chart := range t {
		carr[chart-'a']--
	}
	for _, v := range carr {
		if v != 0 {
			return false
		}
	}
	return true
}

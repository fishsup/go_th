package strtest

import "testing"

/* string与其他语言的差异
1.string是数据类型,不是引用或者指针类型
2.string是只读的byte slice,len函数可以返回它所包含的byte数
3.string的byte数组可以存放任何数据

字符串函数常用包:
1.strings包
2.strconv包
*/

func TestString(t *testing.T) {
	var s string
	t.Log(s)

	s = "test"
	t.Log(s)
	t.Log(len(s))

	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))

	s = "中"
	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF-8 %x", s)
}

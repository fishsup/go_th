package type_test

import (
	"math"
	"testing"
)

type MyInt = int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	var c MyInt

	//不支持隐式类型转换
	// b = a
	b = int64(a)

	//不支持别名的隐式类型转换
	// c = b
	c = MyInt(b)

	t.Log(a, b, c)

	// go 预定义值
	t.Log(math.MaxInt, math.MaxFloat64)
}

/* 与其他语言差异
1.不支持指针运算
2.string是值类型, 其默认的初始化值为空字符串, 而不是nil
*/
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)

	// 不支持指针运算 invalid operation: aPtr + 1 (mismatched types *int and int)
	// aPtr = aPtr + 1
}

func TestString(t *testing.T) {
	//string是值类型, 其默认的初始化值为空字符串, 而不是nil
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}

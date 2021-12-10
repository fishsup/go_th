package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
	1.可以有多个返回值
	2.所有参数都是值传递: slice map channel会有传引用的错觉
	3.函数可以作为变量的值
	4.函数可以作为参数或返回值
*/

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(32)
}

func TestFn(t *testing.T) {
	for i := 0; i < 20; i++ {
		a, _ := returnMultiValues()
		t.Log(a)
	}
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestSlowFunc(t *testing.T) {
	tssf := timeSpent(slowFunc)
	tssf(1)

	timeSpent(slowFunc)(1)
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
}

//defer函数
func TestDefer(t *testing.T) {
	defer func() {
		t.Log("clear Reources")
	}()
	t.Log("started")
	panic("fatal error")
}

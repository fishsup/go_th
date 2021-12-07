package fib

import (
	"testing"
)

//赋值可以进行自动类型推断
func TestFibList(t *testing.T) {

	var (
		a = 1
		b = 1
	)
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		tmp := a
		a = b
		b = tmp + a
	}
}
//在一个赋值语句中可以对多个变量同时进行赋值
func TestEcchange(t *testing.T) {
	a := 1
	b := 2

	//交换两个变量的值
	// tmp := a
	// a = b
	// b = tmp
	a,b = b,a
	t.Log(a, b)

}

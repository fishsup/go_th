package constant_test

import "testing"

//常量自增定义
const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

//常量按位定义
const (
	Readable = 1 << iota
	Writable
	Excutable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
	a := 7
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Excutable == Excutable)
}

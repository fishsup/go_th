package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	// c := [...]int{1, 2, 3, 4}
	d := [...]int{1, 2, 4}
	t.Log(a == b)

	//长度不同的数组比较直接编译错误
	// t.Log(a == c)

	t.Log(a == d)
}

// &^ 按位清零运算符(右边的数对应位上为1,则结果对应位上为0)

const (
	Readable = 1 << iota
	Writable
	Excutable
)

func TestBitClear(t *testing.T) {
	a := 7
	a = a &^ Readable
	t.Log(a)

	a = a &^ Writable
	t.Log(a)

	a = a &^ Excutable
	t.Log(a)
}

package condition_test

import (
	"fmt"
	"testing"
)

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		fmt.Println(a)
	}

	if _, err := SomeFun(); err == nil {
		t.Log("success")
	} else if err != nil {
		t.Log("error")
	}
}

func SomeFun() (int, error) {
	return 1, nil
}

/*
	switch
	1.条件表达式不限制为常数或者整数
	2.单个case中,可以出现多个结果选项,使用逗号分隔
	3.不需要用break来明确退出一个case
	4.可以不设定switch之后的条件表达式, 这种情况下以多个ifelse作用相同
*/
func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("ou")
		case 1, 3:
			t.Log("ji")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchMultiCase1(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("ou")
		case i%2 == 1:
			t.Log("ji")
		default: 
			t.Log("error")
		}
	}
}

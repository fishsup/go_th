package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

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

func slowFunc(op int) int{
	time.Sleep(time.Second*1)
	return op
}

func TestSlowFunc(t *testing.T){
	tssf:=timeSpent(slowFunc)
	tssf(1)

	timeSpent(slowFunc)(1)
}

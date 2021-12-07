package slice

import (
	"reflect"
	"testing"
)

/* 数组vs切片
1.容量是否可伸缩
2.是否可以进行比较
*/
func TestSlice(t *testing.T) {
	var s0 []int
	var s3 [2]int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1), reflect.TypeOf(s1))
	t.Log(len(s3), cap(s3), reflect.TypeOf(s3))

	s4 := make([]int, 3, 5)
	t.Log(len(s4), cap(s4))
	for idx, element := range s4 {
		t.Log(idx, element)
	}
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(summer, len(summer), cap(summer))
	t.Log(Q2)
}

func TestSliceConpare(t *testing.T) {
	// a:=[]int{1,2,3}
	// b:=[]int{2,3,4}
	//不能直接比较slice
	// if a == b {

	// }
}

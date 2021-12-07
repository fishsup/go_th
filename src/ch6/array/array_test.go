package array

import (
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3}
	arr[0] = 5
	t.Log(arr)
	t.Log(arr1)
	t.Log(arr2)
}

func TestArrayTravel(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	len := len(arr)
	for i := 0; i < len/2; i++ {
		arr[i], arr[len-1-i] = arr[len-1-i], arr[i]
	}
	for idx, e := range arr {
		t.Log(idx, e)
	}
}

//数组截取 a[开始索引:结束索引] 包头不包尾
//slice  ![切片内部结构](https://tva1.sinaimg.cn/large/008i3skNgy1gx5eoykop7j32280jy0ug.jpg)
func TestTruncArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	t.Log(a[1:2])
	t.Log(a[1:3])
	t.Log(a[1:4])
	t.Log(a[1:5])
	t.Log(a[1:])
	t.Log(reflect.TypeOf(a[1:]))

	t.Log(reflect.DeepEqual(a[0:], a[0:]))

	var test = a[0:]
	a[0] = 100
	t.Log(test)

	len := len(a)
	for i := 0; i < len/2; i++ {
		a[i], a[len-1-i] = a[len-1-i], a[i]
	}
	t.Log(test)

	test[0] = 101
	test = append(test, 1)

	t.Log(test)
	t.Log(a)

}

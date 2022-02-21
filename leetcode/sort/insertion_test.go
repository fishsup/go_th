package sort

import "testing"

/**
插入排序: 不断地将尚未排好序的数据插入到已经排好序的部分
*/

// 217958  从左到右 从小到大
func sort3(arr []int) {
	len := len(arr)
	for i := 1; i < len; i++ {
		current := arr[i]
		var j int
		for j = i - 1; j >= 0 && arr[j] > current; j-- {
			//元素后移
			arr[j+1] = arr[j]
		}
		arr[j+1] = current
	}
}

func TestSort3(t *testing.T) {
	arr := []int{2, 1, 7, 9, 5, 8}
	sort3(arr)
	t.Log(arr)
}

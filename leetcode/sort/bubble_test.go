package sort

import "testing"

/**
冒泡排序的算法思想
每一轮，从杂乱无章的数组头部开始，每两个元素比较大小并进行交换;
直到这一轮当中最大或最小的元素被放置在数组的尾部；
然后，不断地重复这个过程，直到所有元素都排好位置。
*/

// 217958  从左到右 从小到大

func sort(arr []int) {
	len := len(arr)
	for j := 0; j < len-1; j++ {
		for i := 0; i < len-j-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}

func sort2(arr []int) {
	len := len(arr)
	change := true
	for j := 0; j < len-1 && change; j++ {
		change = false
		for i := 0; i < len-j-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				change = true
			}
		}
	}
}

func TestName(t *testing.T) {
	arr := []int{2, 1, 7, 9, 5, 8}
	sort2(arr)
	t.Log(arr)
}

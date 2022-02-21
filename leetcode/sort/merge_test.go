package sort

import "testing"

/**
// 分治 把一个复杂问题拆分成若干个子问题来求解

1.把数组从中间划分成两个子数组
2.一直递归的把子数组划分成更小的子数组, 直到子数组中只有一个元素
3.依次按照递归的返回顺序, 不断地合并排好序的子数组, 直到把整个数组的顺序排好

T(N)=2*T(N/2)+O(n)
对于规模为N的问题 一共需要进行log(N)层的大小切分 每一层的合并复杂度为O(N)
整体复杂度为O(NlogN)

空间复杂度O(N)
*/
func sortM(arr []int, low int, high int) {
	if low >= high {
		return
	}
	mid := low + (high-low)/2
	sortM(arr, low, mid)
	sortM(arr, mid+1, high)
	merge(arr, low, mid, high)
}

func merge(arr []int, low int, mid int, high int) {
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)
	i := low
	j := low
	k := mid + 1
	for i <= high {
		if j > mid {
			arr[i] = copyArr[k]
			i++
			k++
		} else if k > high {
			arr[i] = copyArr[j]
			i++
			j++
		} else if copyArr[j] > copyArr[k] {
			arr[i] = copyArr[k]
			i++
			k++
		} else {
			arr[i] = copyArr[j]
			i++
			j++
		}
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{2, 1, 7, 9, 5, 8}
	sortM(arr, 0, len(arr)-1)
	t.Log(arr)
}

package sort

import (
	"math/rand"
	"testing"
)

/**
分治
1.选定基准值
2.将数组根据基准值分为两个数组
3.递归1\2步

最优情况下, 基准值将数组均分
T(N)=2*T(N/2)+O(n) 整体复杂度为O(NlogN)
最坏的情况下基准值每次都选为了子数组中的最大或者最小值   复杂度为O(N^2)
空间复杂度O(logn) 取决于递归次数
*/
func qSort(arr []int, idxs int, idxe int) {
	if idxs >= idxe {
		return
	}
	picked := rand.Int()%(idxe-idxs+1) + idxs
	arr[picked], arr[idxs] = arr[idxs], arr[picked]
	d := arr[idxs]
	j := idxs
	for i := idxs + 1; i <= idxe; i++ {
		if arr[i] < d {
			arr[i], arr[j+1] = arr[j+1], arr[i]
			j++
		}
	}
	arr[j], arr[idxs] = arr[idxs], arr[j]
	qSort(arr, idxs, j)
	qSort(arr, j+1, idxe)
}

func TestQsort(t *testing.T) {
	arr := []int{2, 1, 7, 9, 5, 8}
	qSort(arr, 0, len(arr)-1)
	t.Log(arr)
}

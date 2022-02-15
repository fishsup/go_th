package priorityqueue

import (
	"math/rand"
	"testing"
)

func TestHeap(t *testing.T) {
	heap := BinaryIntHeap{elements: [][2]int{}}
	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		heap.push([2]int{a, a})
	}
	for i := 0; i < 10; i++ {
		t.Log(heap.pop())
	}
}

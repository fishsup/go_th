package passing_ref

import "testing"

const NUM_OF_ELEMENTS = 1000

type Content struct {
	Detail [10000]int
}

func withValue(arr [NUM_OF_ELEMENTS]Content) int {
	return 0
}

func withReference(arr *[NUM_OF_ELEMENTS]Content) int {
	return 0
}

func TestFn(t *testing.T) {
	var arr [NUM_OF_ELEMENTS]Content
	withValue(arr)
	withReference(&arr)
}

func BenchmarkPassingArrayWithValue(b *testing.B) {
	var arr [NUM_OF_ELEMENTS]Content
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withValue(arr)
	}
	b.StopTimer()
}

func BenchmarkPassingArrayWithRef(b *testing.B) {
	var arr [NUM_OF_ELEMENTS]Content
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withReference(&arr)
	}
	b.StopTimer()
}

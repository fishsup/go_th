package benchmark

import (
	"bytes"
	"testing"
)

/*
$ go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: go_ph/src/ch26/benchmark
BenchmarkConcatStringByAdd-8            11720929                90.61 ns/op           16 B/op          4 allocs/op
BenchmarkConcatStringByBuffer-8         30474705                38.85 ns/op           64 B/op          1 allocs/op
PASS
ok      go_ph/src/ch26/benchmark        2.477s
*/

func BenchmarkConcatStringByAdd(t *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		ret := ""
		for _, v := range elems {
			ret += v
		}
	}
	t.StopTimer()
}

func BenchmarkConcatStringByBuffer(t *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var buf bytes.Buffer
		for _, v := range elems {
			buf.WriteString(v)
		}
	}
	t.StopTimer()
}

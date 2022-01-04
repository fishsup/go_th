package easyjson

import (
	"encoding/json"
	"testing"
)

var str = `{
	"basic_info": {
		"name": "etest",
		"age": 19
	},
	"job_info": {
		"skills": [
			"1",
			"2"
		]
	}
}`

/* 
$ go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: go_ph/src/ch31/easy_json
BenchmarkEmbeddedJson-8           569312              1866 ns/op             496 B/op         11 allocs/op
BenchmarkEasyJson-8              1819621               660.8 ns/op           240 B/op          2 allocs/op
PASS
ok      go_ph/src/ch31/easy_json        3.038s */

func BenchmarkEmbeddedJson(b *testing.B) {
	b.ResetTimer()
	e := new(EmployeeEmbedded)
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(str), e)
		if err != nil {
			b.Error(err)
		}
		if _, err := json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEasyJson(b *testing.B) {
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		err := e.UnmarshalJSON([]byte(str))
		if err != nil {
			b.Error(err)
		}
		if _, err := e.MarshalJSON(); err != nil {
			b.Error(err)
		}
	}
}

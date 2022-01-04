package easyjson

import (
	"testing"
)

func TestEmbeddedJson(t *testing.T) {
	b := BasicInfo{Age: 19, Name: "etest"}
	basic_info_json, _ := b.MarshalJSON()
	t.Log(string(basic_info_json))

	j := JobInfo{Skills: []string{"1", "2"}}
	job_info_json, _ := j.MarshalJSON()
	t.Log(string(job_info_json))

	e := Employee{BasicInfo: b, JobInfo: j}
	if v, err := e.MarshalJSON(); err == nil {
		t.Log(string(v))
	} else {
		t.Error(err)
	}

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
	ue := new(Employee)
	ue.UnmarshalJSON([]byte(str))
	t.Log(ue)
}

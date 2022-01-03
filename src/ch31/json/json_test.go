package json

import (
	"encoding/json"
	"testing"
)

/* go内置 利用反射实现,
通过FieldTag来标识对应的json值 */
type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type JobInfo struct {
	Skills []string `json:"skills"`
}

type Employee struct {
	BasicInfo *BasicInfo `json:"basic_info"`
	JobInfo   *JobInfo   `json:"job_info"`
}

func TestEmbeddedJson(t *testing.T) {
	b := BasicInfo{Age: 19, Name: "etest"}
	//json.Marshal返回[]byte与err
	basic_info_json, _ := json.Marshal(b)
	t.Log(string(basic_info_json))

	j := JobInfo{Skills: []string{"1", "2"}}
	job_info_json, _ := json.Marshal(j)
	t.Log(string(job_info_json))

	e := Employee{BasicInfo: &b, JobInfo: &j}
	if v, err := json.Marshal(e); err == nil {
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
	json.Unmarshal([]byte(str), ue)
	t.Log(ue.BasicInfo)
	t.Log(ue.JobInfo)
}

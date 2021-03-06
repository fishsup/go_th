package easyjson

import (
	"encoding/json"
	"testing"
)

/* go内置 利用反射实现,
通过FieldTag来标识对应的json值 */
type BasicInfoEmbedded struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type JobInfoEmbedded struct {
	Skills []string `json:"skills"`
}

type EmployeeEmbedded struct {
	BasicInfoEmbedded  `json:"basic_info"`
	JobInfoEmbedded      `json:"job_info"`
}

func TestEmbeddedJson1(t *testing.T) {
	b := BasicInfoEmbedded{Age: 19, Name: "etest"}
	//json.Marshal返回[]byte与err
	basic_info_json, _ := json.Marshal(b)
	t.Log(string(basic_info_json))

	j := JobInfoEmbedded{Skills: []string{"1", "2"}}
	job_info_json, _ := json.Marshal(j)
	t.Log(string(job_info_json))

	e := EmployeeEmbedded{BasicInfoEmbedded: b, JobInfoEmbedded: j}
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
	ue := new(EmployeeEmbedded)
	json.Unmarshal([]byte(str), ue)
	t.Log(ue.BasicInfoEmbedded)
	t.Log(ue.JobInfoEmbedded)
}

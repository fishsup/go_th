package roa

import "log"

type Response struct {
	Success bool
	Data    string
	Msg     string
}

func ResponseOk(json string) []byte {
	r := Response{Success: true, Data: json}
	if v, err := r.MarshalJSON(); err != nil {
		log.Fatal(err)
		return nil
	} else {
		return v
	}
}

func ResponseFail(errorMsg string) []byte {
	r := Response{Success: false, Data: "", Msg: errorMsg}
	if v, err := r.MarshalJSON(); err != nil {
		log.Fatal(err)
		return nil
	} else {
		return v
	}
}

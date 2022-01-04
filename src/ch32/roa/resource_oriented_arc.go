package main

import (
	"fmt"
	roa "go_ph/src/ch32/roa_struct"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var employeeMap map[string]*roa.Employee

func init() {
	employeeMap = map[string]*roa.Employee{}
	employeeMap["Mike"] = &roa.Employee{Name: "Mike", Age: 1, Id: "a"}
	employeeMap["Rose"] = &roa.Employee{Name: "Rose", Age: 2, Id: "b"}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func GetEmployeeByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	qName := ps.ByName("name")
	fmt.Println(qName)
	if v, ok := employeeMap[qName]; ok {
		if b, toJsonErr := v.MarshalJSON(); toJsonErr == nil {
			w.Write(roa.ResponseOk(string(b)))
		}
		return
	} else {
		w.Write(roa.ResponseFail("employee not found"))
		return
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/employee/:name", GetEmployeeByName)

	log.Fatal(http.ListenAndServe(":9850", router))
}

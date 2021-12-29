package reflecttest

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type T struct {
	Name    string
	Numbers []float64
}

func TestDeepEqual(t *testing.T) {

	a := map[int]string{1: "1", 2: "2", 3: "3"}
	b := map[int]string{1: "1", 2: "2", 3: "3"}

	//invalid operation: a == b (map can only be compared to nil)
	// fmt.Println(a == b)
	fmt.Println(reflect.DeepEqual(a, b))

	n1 := []float64{1.13459, 2.29343, 3.010100010}
	n2 := []float64{1.13459, 2.29343, 3.010100010}
	//invalid operation: n1 == n2 (slice can only be compared to nil)
	// fmt.Println(n1 == n2)
	fmt.Println(reflect.DeepEqual(n1, n2))
}

/* 1.提高程序灵活性
2.降低可读性
3.降低性能 */

type Employee struct {
	EmployeeId string
	Name       string `format:"normal"`
	Age        int
}

type Customer struct {
	CookieId string
	Name     string `format:"normal"`
	Age      int
}

func TestFillNameAndAge(t *testing.T) {
	smap := map[string]interface{}{"Name": "testName", "Age": 1}
	e := &Employee{}
	if err := FillNameAndAge(e, smap); err == nil {
		t.Log(e)
	} else {
		t.Error(err)
	}
	c := new(Customer)
	if err := FillNameAndAge(c, smap); err == nil {
		t.Log(e)
	} else {
		t.Error(err)
	}

}

func FillNameAndAge(c interface{}, smap map[string]interface{}) error {
	objType := reflect.TypeOf(c)

	fmt.Println(objType.Kind())
	fmt.Println(objType.Elem().Kind())

	if objType.Kind() != reflect.Ptr ||
		objType.Elem().Kind() != reflect.Struct {
		return errors.New("must be ptr")
	}
	for k, v := range smap {
		if field, ok := objType.Elem().FieldByName(k); !ok {
			continue
		} else {
			if field.Type == reflect.TypeOf(v) {
				reflect.ValueOf(c).Elem().FieldByName(k).Set(reflect.ValueOf(v))
			}
		}
	}
	return nil
}

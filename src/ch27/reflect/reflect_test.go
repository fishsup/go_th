package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

/* reflect.TypeOf 返回类型 reflect.Type
reflect.ValueOf 返回值 reflect.Value
可以从reflect.Value获取类型 通过kind来判断类型
*/
func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("FLOAT")
	case reflect.Int, reflect.Int16, reflect.Int64:
		fmt.Println("INT")
	default:
		fmt.Println(t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(&f)
	CheckType(f)

	e := &Employee{"1", "name", 3}
	CheckType(e)
	CheckType(*e)

	t.Log(reflect.ValueOf(e))
	t.Log(reflect.ValueOf(*e))
}

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

type Employee struct {
	EmployeeId string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newAge int) {
	e.Age = newAge
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "name", 3}
	t.Logf("name :value(%[1]v), type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); ok {
		t.Log("TAG:FORMAT", nameField.Tag.Get("format"))
	} else {
		t.Error("failed to get field name")
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("updated age", e)
}

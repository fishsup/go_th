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
		fmt.Println("NO CARE")
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(&f) 
	CheckType(f) 
}

func TestTypeAndValue(t *testing.T){
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

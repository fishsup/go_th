package polymorphic

import (
	"fmt"
	"testing"
)

/* 空接口与断言
1.空接口可以标识任何类型, 类似c\c++中的void* 和java中的Object
2.通过断言来讲空接口转换为指定类型
	v,ok := p.(int) // ok=true时转化成功
*/
func DoSomething(p interface{}) {
	// if v, ok := p.(int); ok {
	// 	fmt.Println("int", v)
	// 	return
	// }
	// if v, ok := p.(string); ok {
	// 	fmt.Println("string", v)
	// 	return
	// }
	// fmt.Println("unknown type")

	switch v := p.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unknown type")
	}

}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}

/* go接口最佳实践 */
// 1.倾向于使用小的接口定义, 很多接口只包含一个方法
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Writer(p []byte) (n int, err error)
}

// 2.较大的接口定义, 可以有多个小接口定义组合而成
type ReaderWriter interface {
	Reader
	Writer
}

// 3.只依赖与必要功能的最小接口
func StoreData(reader Reader) {

}

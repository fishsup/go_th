package interface_test

import (
	"testing"
)

/* dock type
1.接口为非入侵性, 实现不依赖于接口定义
2.接口的定义可以包含在接口使用者包内
*/
type Programmer interface {
	WriteHelloworld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloworld() string {
	return "fmt.Println(\"hello world\")"
}

func TestClient(t *testing.T) {
	// p -> 接口变量  类型为GoProgrammer 数据为GoProgrammer的实例
	var p Programmer = new(GoProgrammer)
	t.Log(p.WriteHelloworld())
}

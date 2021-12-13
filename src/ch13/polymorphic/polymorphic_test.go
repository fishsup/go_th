package polymorphic

import (
	"fmt"
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

type JavaProgrammer struct {
}

func (g *GoProgrammer) WriteHelloworld() string {
	return "fmt.Println(\"hello world\")"
}

func (g *JavaProgrammer) WriteHelloworld() string {
	return "System.out.Println(\"hello world\")"
}

// Programmer是一个接口, 对应的传值类型只能是一个指针, 不能是实例
func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloworld())
}

func TestClient(t *testing.T) {
	// p -> 接口变量  类型为GoProgrammer 数据为GoProgrammer的实例
	goProgrammer := new(GoProgrammer)
	// Programmer是一个接口, 对应的传值类型只能是一个指针, 不能是实例
	// JavaProgrammer := new(JavaProgrammer
	JavaProgrammer := &JavaProgrammer{}
	writeFirstProgram(goProgrammer)
	writeFirstProgram(JavaProgrammer)
}

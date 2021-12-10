package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"test", "test", 1}
	t.Log(e)

	e1 := Employee{Id: "test"}
	t.Log(e1)

	e2 := new(Employee)
	e2.Id = "et"
	e2.Name = "dsf"
	e2.Age = 20
	t.Log(e2)
}

// 行为定义
// 第一种方式在实例对应方法被调用时, 实例的成员会进行值复)
func (e Employee) String() string {
	fmt.Printf("String is %x   %x \n", unsafe.Pointer(&e), unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

//通常情况下为了避免内存拷贝我们使用第二种方式(实例指针是不同的)
func (e *Employee) PointerString() string {
	fmt.Printf("PointerString is %x   %x \n", unsafe.Pointer(&e), unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"id", "name", 12}
	fmt.Printf("Address is %x   %x \n", unsafe.Pointer(&e), unsafe.Pointer(&e.Name))
	t.Log(e.String())
	t.Log(e.PointerString())
}

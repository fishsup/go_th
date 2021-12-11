package extension

import (
	"fmt"
	"testing"
)

type Pet struct{}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

// type Dog struct {
// 	p *Pet
// }

// 匿名嵌套类型 类似继承,但是重载后的方法对于继承的父类的其他方法是无法感知的
type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Println("wang!")
}

// func (d *Dog) SpeakTo(host string) {
// 	d.p.SpeakTo(host)
// }

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speak()
	dog.SpeakTo("test")
}

package map_test

import "testing"

/*
map的value可以是一个方法
与go的Dock type接口方式一起,可以方便的实现单一方法对象的工厂模式
*/
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2))
	t.Log(m[2](2))
	t.Log(m[3](2))
}

/*
实现set
go内置的集合中没有set实现, 可以map[type]bool
1.元素的唯一性
2.基本操作(添加元素\判断元素是否存在\删除元素\元素个数)
*/
func TestMapforSet(t *testing.T) {
	myset := map[int]bool{}
	myset[1] = true
	for i := 0; i < 10; i++ {
		if myset[i] {
			t.Logf("%d is existing", i)
		} else {
			t.Logf("%d is not existing", i)
		}
	}
	t.Log(len(myset))
	delete(myset, 1)
	t.Log(len(myset))
}

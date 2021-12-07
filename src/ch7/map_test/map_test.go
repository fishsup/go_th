package map_test

import "testing"

func TestMapInit(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(m1[2])
	t.Log("len m1:", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Log("len m2:", len(m2))
	m3 := make(map[int]int, 10)
	t.Log("len m3:", len(m3))
}

/* 当访问的key不存在时, 仍然会返回0, 不能通过nil来判断元素是否存在 */
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("key 3 is existing %d", v)
	} else {
		t.Log("key 3 is not existing")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}

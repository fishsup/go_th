package _map

import "sync"

type RWLockMap struct {
	m    map[interface{}]interface{}
	lock sync.RWMutex
}

func (m *RWLockMap) Set(key interface{}, val interface{}) {
	m.lock.Lock()
	m.m[key] = val
	m.lock.Unlock()
}
func (m *RWLockMap) Get(key interface{}) (interface{}, bool) {
	m.lock.RLock()
	v, ok := m.m[key]
	m.lock.RUnlock()
	return v, ok
}
func (m *RWLockMap) Del(key interface{}) {
	m.lock.Lock()
	delete(m.m, key)
	m.lock.Unlock()
}

func NewRwLockMap() *RWLockMap {
	m := make(map[interface{}]interface{}, 0)
	return &RWLockMap{m: m}
}

package _map

import "sync"

type SyncMap struct {
	m sync.Map
}

func (m *SyncMap) Set(key interface{}, val interface{}) {
	m.m.Store(key, val)
}
func (m *SyncMap) Get(key interface{}) (interface{}, bool) {
	return m.m.Load(key)
}
func (m *SyncMap) Del(key interface{}) {
	m.m.Delete(key)
}

func NewSyncMap() *SyncMap {
	return &SyncMap{}
}

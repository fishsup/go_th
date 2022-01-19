package _map

import (
	"github.com/easierway/concurrent_map"
)

type ConcurrentMap struct {
	cm *concurrent_map.ConcurrentMap
}

func (m *ConcurrentMap) Set(key interface{}, val interface{}) {
	m.cm.Set(concurrent_map.StrKey(key.(string)), val)
}
func (m *ConcurrentMap) Get(key interface{}) (interface{}, bool) {
	return m.cm.Get(concurrent_map.StrKey(key.(string)))
}
func (m *ConcurrentMap) Del(key interface{}) {
	m.cm.Del(concurrent_map.StrKey(key.(string)))
}

func NewConcurrentMap(numOfPartitions int) *ConcurrentMap {
	cm := concurrent_map.CreateConcurrentMap(numOfPartitions)
	return &ConcurrentMap{cm: cm}
}

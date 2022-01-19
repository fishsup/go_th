package _map

import (
	"strconv"
	"sync"
	"testing"
)

const (
	NumOfReader = 200
	NumOfWriter = 20
)

type Map interface {
	Set(key interface{}, val interface{})
	Get(key interface{}) (interface{}, bool)
	Del(key interface{})
}

func benchmarkMap(b *testing.B, hm Map) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for i := 0; i < NumOfWriter; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100; i++ {
					hm.Set(strconv.Itoa(i), i*i)
					hm.Set(strconv.Itoa(i), i*i)
					hm.Del(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
		for i := 0; i < NumOfReader; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100; i++ {
					hm.Get(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkMap(b *testing.B) {
	b.Run("map with RWLOCK", func(b *testing.B) {
		hm := NewRwLockMap()
		benchmarkMap(b, hm)
	})
	b.Run("sync.Map", func(b *testing.B) {
		hm := NewSyncMap()
		benchmarkMap(b, hm)
	})
	b.Run("concurrentmap", func(b *testing.B) {
		hm := NewConcurrentMap(199)
		benchmarkMap(b, hm)
	})
}

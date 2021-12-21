package singletest

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct{}

var singletonInstance *Singleton
var once sync.Once

func GetSingletonInstance() *Singleton {
	once.Do(func() {
		fmt.Println("create obj")
		singletonInstance = new(Singleton)
	})
	return singletonInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonInstance()
			fmt.Printf("%d\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}

package sharemem

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	var lock sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				lock.Unlock()
				wg.Done()
			}()
			lock.Lock()
			counter++
		}()
	}
	// time.Sleep(1 * time.Second)
	wg.Wait()
	t.Log(counter)
}

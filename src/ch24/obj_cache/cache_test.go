package objcache

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

/* sync.Pool
1.适合于通过复用,降低复杂对象的创建和GC代价
2.协程安全, 会有锁的开销
3.生命周期收到GC影响 

获取对象:
1.尝试从私有对象(协程安全)获取
2.私有对象不存在, 尝试从当前Processor的共享池(协程不安全的)获取
3.如果当前Processor的共享池是空的, 尝试从其他Processor的共享池获取
4.如果所有自持都是空的,最后就用用户的指定New函数创建一个新的对象返回

放回对象:
1.如果私有对象不存在则保存为私有对象
2.如果私有对象存在,放入当前Processor子池的共享池中

*/

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create new obj")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)

	runtime.GC() //gc会清除sync.pool中缓存的对象

	v1, _ := pool.Get().(int)
	fmt.Println(v1)


	// get之后私有对象就没了, 下次获取前不put还是会创建
	v2, _ := pool.Get().(int)
	fmt.Println(v2)

	v3, _ := pool.Get().(int)
	fmt.Println(v3)

}

func TestSyncPoolInMultiGoroutine(t *testing.T){
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create new obj")
			return 10
		},
	}

	pool.Put(1)
	pool.Put(2)
	pool.Put(3)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()

}

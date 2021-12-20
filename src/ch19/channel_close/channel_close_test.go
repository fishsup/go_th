package channelclose

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		// ch <- 11   向关闭的通道继续发送消息会导致 panic: send on closed channel
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 11; i++ {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				fmt.Println("channel closed")
				break
			}
		}
		wg.Done()
	}()
}

func TestData(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
	fmt.Println("done")
}

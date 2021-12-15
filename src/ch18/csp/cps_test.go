package csp

import (
	"fmt"
	"testing"
	"time"
)
/* Do not communicate by sharing memory; instead, share memory by communicating.
Communicating Sequential Processes
CSP vs actor model

1.actor直接通讯, csp通过channel
2.go的channel有容量限制, goroutine会主动地从channel中拉取消息
  而erlang actor模式中的mailbox是容量无限的, 接受进程也是被动的接受消息

*/

func TestCsp(t *testing.T) {
	fmt.Println(service())
	fmt.Println(otherTask())
}

func TestAsyncCsp(t *testing.T) {
	retCh := asyncService()
	fmt.Println(otherTask())
	fmt.Println(<-retCh)
	
}

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "service done"
}

func otherTask() string {
	time.Sleep(time.Millisecond * 100)
	return "other task done"
}

func asyncService() chan string {
	// retCh := make(chan string)
	// buffered channel
	retCh := make(chan string,1)
	go func() {
		ret := service()
		fmt.Println("asyncService return result")
		retCh <- ret
		fmt.Println("asyncService exit")
	}()
	return retCh
}

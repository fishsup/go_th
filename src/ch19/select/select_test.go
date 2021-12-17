package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "service done"
}

func otherTask() string {
	time.Sleep(time.Millisecond * 100)
	return "other task done"
}

func asyncService() chan string {
	// retCh := make(chan string)
	// buffered channel
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("asyncService return result")
		retCh <- ret
		fmt.Println("asyncService exit")
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-asyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Log("time out")
	}
}

package poo_test

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

/* 使用buffered channel实现对象池 */

type ReusableObj struct{}

type ReusableObjPool struct {
	bufferchan chan *ReusableObj
}

func NewPool(numOfObj int) *ReusableObjPool {
	objpool := ReusableObjPool{}
	objpool.bufferchan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objpool.bufferchan <- &ReusableObj{}
	}
	return &objpool
}

func (p *ReusableObjPool) GetObjWithTimeOut(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufferchan:
		return ret, nil
	case <-time.After(timeout):
		// quick failure better than slow resp
		return nil, errors.New("get obj timeout")
	}
}

func (p *ReusableObjPool) GetObj() *ReusableObj {
	return <-p.bufferchan
}

func (p *ReusableObjPool) releaseObj(obj *ReusableObj) error {
	select {
	case p.bufferchan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestObjPool(t *testing.T) {
	pool := NewPool(10)
	//对象池是满的, 再放入会直接异常
	// if err := pool.releaseObj(&ReusableObj{}); err != nil {
	// 	t.Error(err)
	// }

	for i := 0; i < 10; i++ {
		if v, err := pool.GetObjWithTimeOut(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			// if err := pool.releaseObj(v); err != nil {
			// 	t.Error(err)
			// }
		}
	}
	fmt.Println("DONE")

}

package ca_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func is_cancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if is_cancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "cancelled")
		}(i, ctx)
	}
	// cancel_1(cancelChan)
	cancel()
	time.Sleep(time.Second * 1)
}

package test

import (
	"fmt"
	"sync/atomic"
	"time"
)

func TestSum() int {
	var ops uint64
	go func() {
		for i := 0; i < 1000; i++ {
			atomic.AddUint64(&ops, 1)
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			atomic.AddUint64(&ops, 1)
			time.Sleep(time.Millisecond)
		}
	}()
	fmt.Println(ops)
	time.Sleep(time.Second * 1)
	return int(ops)
}

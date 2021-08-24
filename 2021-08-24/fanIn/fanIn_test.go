package fanIn

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
		fmt.Println("任务进行中，当前协程数:", after, runtime.NumGoroutine())
	}()
	return c
}

// 扇入测试
func Test_fanInReflect(t *testing.T) {
	start := time.Now()

	<-fanInReflect(
		sig(2*time.Second),
		sig(1*time.Second),
		sig(6*time.Second),
		sig(4*time.Second),
		sig(5*time.Second),
		sig(1*time.Minute),
		sig(1*time.Minute),
		sig(1*time.Minute),
		sig(1*time.Minute),
		sig(1*time.Minute),
		sig(1*time.Minute),
		sig(1*time.Minute),
		sig(1*time.Minute),
		sig(1*time.Minute),
	)

	fmt.Printf("fanInReflect done after %v", time.Since(start))
}

package _chan

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
		fmt.Println("任务进行中，当前协程数:", runtime.NumGoroutine())
	}()
	return c
}

// 测试orReflect
func Test_orReflect(t *testing.T) {
	start := time.Now()

	<-orReflect(
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

	fmt.Printf("orReflect done after %v", time.Since(start))
}

// 测试orDirect
func Test_orDirect(t *testing.T) {
	start := time.Now()

	<-orDirect(
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

	fmt.Printf("orDirect done after %v", time.Since(start))
}

// 测试orBinary
func Test_orBinary(t *testing.T) {
	start := time.Now()

	<-orBinary(
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

	fmt.Printf("orBinary done after %v", time.Since(start))
}

package mapChan

import (
	"fmt"
	"testing"
	"time"
)

// 测试 map
func Test_mapChan(t *testing.T) {
	for v := range mapChan(func() (out chan interface{}) {
		out = make(chan interface{})
		go func() {
			for {
				select {
				case t := <-time.NewTicker(time.Second).C:
					out <- t.Second()
				}
			}
		}()
		return out
	}(), func(i interface{}) interface{} {
		return fmt.Sprintf("值为：%v", i)
	}) {
		fmt.Println(v)
	}
}

// 测试reduce
func Test_reduceChan(t *testing.T) {

	in := make(chan interface{})
	go func() {
		defer close(in)
		in <- 5
		in <- 6
		in <- 7
		in <- 8
		in <- 9
	}()

	v := reduceChan(in, func(r, v interface{}) interface{} {
		return fmt.Sprintf("r:%v,v:%v", r, v)
	})
	fmt.Println(v)
}

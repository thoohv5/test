package stream

import (
	"fmt"
	"testing"
	"time"
)

// 测试 Stream模式
func Test_takeN(t *testing.T) {

	done := make(chan struct{})
	for c := range takeN(done, func() (out chan interface{}) {
		out = make(chan interface{})
		go func() {
			for {
				select {
				case <-time.NewTicker(time.Second).C:
					out <- struct{}{}
				}
			}
		}()
		return out
	}(), 10) {
		fmt.Println(c)
	}

}

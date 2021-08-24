package fanOut

import (
	"fmt"
	"testing"
	"time"
)

// 测试扇出
func Test_fanOut(t *testing.T) {

	out1 := make(chan interface{})
	out2 := make(chan interface{})
	out3 := make(chan interface{})
	out4 := make(chan interface{})

	fanOut(func() (out chan interface{}) {
		out = make(chan interface{})
		go func() {
			for {
				select {
				case <-time.NewTicker(time.Second).C:
					fmt.Println("time ticker c")
					out <- struct{}{}
				}
			}
		}()
		return
	}(), []chan interface{}{out1, out2, out3, out4}, false)

	for {
		select {
		case c := <-out1:
			fmt.Printf("out1:%v\n", c)
		case c := <-out2:
			fmt.Printf("out2:%v\n", c)
		case c := <-out3:
			fmt.Printf("out3:%v\n", c)
		case c := <-out4:
			fmt.Printf("out4:%v\n", c)

		}
	}

}

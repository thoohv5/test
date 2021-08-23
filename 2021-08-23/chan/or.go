package _chan

import (
	"reflect"
)

// https://zhuanlan.zhihu.com/p/367824324

// Or-Done 模式是在多个任务场景，只要有一个任务执行完，就会知道这个完成信号
func orReflect(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		reflect.Select(cases)
	}()

	return orDone
}

func orDirect(channels ...<-chan interface{}) <-chan interface{} { // 1

	switch len(channels) {
	case 0: // 2
		return nil
	case 1: // 3
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() { // 4
		defer close(orDone)

		switch len(channels) {
		case 2: // 5
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default: // 6
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-orDirect(append(channels[3:], orDone)...): // 6
			}
		}
	}()
	return orDone
}

func orBinary(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		case 3:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			}
		default:
			m := len(channels) / 2
			select {
			case <-orBinary(append(channels[:m:m], orDone)...): // must append orDone to avoid leak!!!!
			case <-orBinary(append(channels[m:], orDone)...):
			}
		}
	}()

	return orDone
}

package stream

// Stream 模式就是把Channel看作流（Stream），提供跳过几个元素，或者只取其中几个元素的方法
func takeN(done <-chan struct{}, values <-chan interface{}, num int) <-chan interface{} {
	sOut := make(chan interface{})
	go func() {
		defer close(sOut)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case v := <-values:
				sOut <- v
			}
		}
	}()
	return sOut
}

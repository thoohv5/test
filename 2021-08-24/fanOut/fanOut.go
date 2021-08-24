package fanOut

// 扇出模式
func fanOut(ch <-chan interface{}, out []chan interface{}, async bool) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		for v := range ch {
			for i := 0; i < len(out); i++ {
				if async {
					go func(i int, v interface{}) {
						out[i] <- v
					}(i, v)
				} else {
					out[i] <- v
				}
			}
		}
	}()
}

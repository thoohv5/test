package mapChan

// map 多数据，单操作
func mapChan(in <-chan interface{}, fn func(interface{}) interface{}) <-chan interface{} {
	out := make(chan interface{})
	if in == nil {
		close(out)
		return out
	}

	go func() {
		defer close(out)
		for v := range in {
			out <- fn(v)
		}
	}()

	return out
}

// reduce 单数据，多操作
func reduceChan(in <-chan interface{}, fn func(r, v interface{}) interface{}) interface{} {
	if in == nil {
		return nil
	}
	val := <-in
	for v := range in {
		val = fn(val, v)
	}

	return val
}

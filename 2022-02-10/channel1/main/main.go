package main

func main() {
	var ch = make(chan struct{})
	ch = nil
	// close(ch)

	<-ch
}

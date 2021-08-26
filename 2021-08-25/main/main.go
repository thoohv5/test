package main

import "fmt"

// go build -o main -ldflags=-compressdwarf=false -gcflags="-N -l" main.go
func main() {
	a := 1
	b := 2
	fmt.Println(a + b)
}

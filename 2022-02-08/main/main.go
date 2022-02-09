package main

import "fmt"

// go build -gcflags -S main.go
// GOSSAFUNC=main go build main.go
func main() {
	a := [3]int{1, 2, 3}
	fmt.Println(a)
}

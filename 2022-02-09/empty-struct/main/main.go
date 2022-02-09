package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Sizeof(struct {
		m struct{} // 0
		n int8     // 1
	}{})) // 1
	fmt.Println(unsafe.Sizeof(struct {
		n int8     // 1
		m struct{} // 0
	}{})) // 2
}

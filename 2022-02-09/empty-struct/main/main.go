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

	fmt.Println(unsafe.Sizeof(struct {
		m struct{} // 0
		n int64    // 8
	}{})) // 8
	fmt.Println(unsafe.Sizeof(struct {
		n int64    // 8
		m struct{} // 0
	}{})) // 16

	fmt.Println(unsafe.Sizeof(struct {
		m struct{} // 0
		a int8     // 1
		n int64    // 8
	}{})) // 16
	fmt.Println(unsafe.Sizeof(struct {
		n int64    // 8
		a int64    // 1
		m struct{} // 0
	}{})) // 16
}

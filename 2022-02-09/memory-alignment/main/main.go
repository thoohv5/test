package main

import (
	"fmt"
	"unsafe"
)

type Ag struct {
	arr [2]int8  // 2
	b1  bool     // 1
	sl  []int16  // 24
	ptr *int64   // 8
	st  struct { // 16
		str string // 16
	}
	m map[string]int16 // 8
	i interface{}      // 16
}

func main() {
	var a Ag
	fmt.Println(unsafe.Sizeof(a), unsafe.Alignof(a))

	var i int
	fmt.Println(unsafe.Sizeof(i), unsafe.Alignof(i))
}

package main

import (
	"fmt"
)

func main() {

	ss := []int{1, 2, 3}
	fmt.Println(ss, ss[0:0], append(ss[:0], ss[1:]...))

}

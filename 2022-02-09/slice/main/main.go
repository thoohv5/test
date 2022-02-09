package main

import "fmt"

// go tool compile -N -l -S main.go
func main() {

	// var s []int64
	//
	// // panic: runtime error: index out of range [0] with length 0
	// fmt.Println(s[0])

	slice := make([]int, 5, 10) // 长度为5，容量为10
	slice[2] = 2                // 索引为2的元素赋值为2
	fmt.Println(slice)

}

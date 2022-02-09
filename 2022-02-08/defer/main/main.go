package main

import "fmt"

func main() {
	a := make([]int, 0)
	b := make([]interface{}, 0)
	var c []string
	d := [...]string{}
	f := [0]string{}
	fmt.Printf("%p, %p, %p, %p, %p", a, b, c, &d, &f)
}

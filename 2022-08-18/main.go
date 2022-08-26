package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	// …
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}

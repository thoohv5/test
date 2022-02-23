package main

import "fmt"

func main() {
	deferLoop()
}

func deferLoop() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println("defer loop", i)
		}()
	}

	return

}

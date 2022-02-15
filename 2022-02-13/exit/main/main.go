package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	go func() {

		defer func() {
			fmt.Println("defer")
		}()

		os.Exit(0)

	}()

	time.Sleep(5 * time.Second)
}

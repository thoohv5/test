package main

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

func main() {
	r, w := io.Pipe()

	go func() {
		writer := bufio.NewWriter(w)
		for {
			time.Sleep(time.Second)
			writer.WriteString(time.Now().String() + "\n")
			// writer.Flush()
		}
	}()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now())
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				time.Sleep(time.Second)
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(4000 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println(time.Now())
	fmt.Println("Ticker stopped")
}

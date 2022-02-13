package main

import (
	"fmt"
	"sync"
)

// 锁实现方式
func main() {
	var count int64
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			mu.Lock()
			count = count + 1
			mu.Unlock()
		}(&wg)
	}
	wg.Wait()

	// count = 10000
	fmt.Println("count = ", count)
}

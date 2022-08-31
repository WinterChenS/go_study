package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var rwMutex sync.RWMutex
	wg := sync.WaitGroup{}

	Data := 0
	wg.Add(20)

	for i := 1; i <= 10; i++ {
		go func(t int) {
			fmt.Printf("Locking (G%d)\n", t)
			rwMutex.RLock()
			defer rwMutex.RUnlock()
			fmt.Printf("Read DataL %v\n", Data)
			wg.Done()
			time.Sleep(time.Second * 2)
		}(i)

		go func(t int) {
			fmt.Printf("Locking (G%d)\n", t)
			rwMutex.Lock()
			defer rwMutex.Unlock()
			Data += t
			fmt.Printf("Write Data %v %d\n", Data, t)
			wg.Done()
			time.Sleep(time.Second * 2)
		}(i)
	}
	wg.Wait()
}

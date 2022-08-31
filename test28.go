package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	wg := sync.WaitGroup{}

	fmt.Println("Locking (G0)")
	mutex.Lock()
	fmt.Println("Locked (G0)")

	wg.Add(3)

	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Printf("Locking (G%d)\n", i)
			mutex.Lock()
			fmt.Printf("Locked (G%d)\n", i)
			time.Sleep(time.Second * 2)
			mutex.Unlock()
			wg.Done()
		}(i)
	}

	time.Sleep(time.Second * 5)
	fmt.Println("Unlocking (G0)")
	mutex.Unlock()
	fmt.Println("Unlocked (G0)")

	wg.Wait()

}


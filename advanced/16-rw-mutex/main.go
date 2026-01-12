package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	rwmu    sync.RWMutex
	counter int
)

func readCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	defer rwmu.RUnlock()
	rwmu.RLock()
	fmt.Printf("Value of counter is: %v\n", counter)
}

func writeCounter(wg *sync.WaitGroup, value int) {
	defer wg.Done()
	defer rwmu.Unlock()
	rwmu.Lock()
	counter = value
}

func main() {

	var wg sync.WaitGroup

	for range 5 {
		wg.Add(1)
		go readCounter(&wg)
	}

	wg.Add(1)
	time.Sleep(3 * time.Microsecond)
	go writeCounter(&wg, 18)

	wg.Wait()

}

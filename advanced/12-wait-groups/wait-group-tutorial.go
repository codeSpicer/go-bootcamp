package main

import (
	"fmt"
	"sync"
	"time"
)

func worker1(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func BasicExample() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)	// always add counter to waitgroup outside go routine
		go worker1(i, &wg)
	}

	wg.Wait()
	fmt.Println("All Go routines finished")

}

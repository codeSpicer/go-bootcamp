package main

import (
	"fmt"
	"sync"
	"time"
)

const bufferSize = 5

type buffer struct {
	items []int
	mu    sync.Mutex
	cond  *sync.Cond
}

func newBuffer(size int) *buffer {
	b := &buffer{items: make([]int, 0, size)}
	// sync.Cond must be initialized with the same mutex used for locking
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *buffer) produce(item int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Wait() blocks if the buffer is full. 
	// It automatically unlocks b.mu while waiting and relocks it when woken up.
	for len(b.items) == bufferSize {
		b.cond.Wait()
	}

	b.items = append(b.items, item)
	fmt.Println("Produced: ", item)

	// Signal one waiting consumer that there is now an item available
	b.cond.Signal()
}

func (b *buffer) consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Wait() blocks if the buffer is empty.
	for len(b.items) == 0 {
		b.cond.Wait()
	}

	item := b.items[0]
	b.items = b.items[1:]
	fmt.Println("Consumed item:", item)

	// Signal one waiting producer that there is now space available
	b.cond.Signal()
	return item
}

func producer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		b.produce(i + 100)
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 10 {
		b.consume()
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	buff := newBuffer(bufferSize)
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(buff, &wg)
	go consumer(buff, &wg)

	wg.Wait()
	fmt.Println("Done.")
}
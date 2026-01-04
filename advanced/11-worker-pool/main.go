package main

import (
	"fmt"
	"time"
)

func worker(id int, tasks <-chan int, results chan<- int) {

	for task := range tasks {
		fmt.Printf("Worker %d is processing task %d...\n", id, task)
		time.Sleep(1 * time.Second)
		results <- task * 2
	}
}

func main() {

	numWorkers := 3
	numJobs := 10

	tasks := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, tasks, results)
	}

	for i := range numJobs {
		tasks <- i
	}
	close(tasks)

	for range numJobs {
		result := <-results
		fmt.Println("Results: ", result)
	}

}

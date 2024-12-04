package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, count *int, barrier chan bool, mutex *sync.Mutex) {
	defer wg.Done() // ChatGpt gave me this idea, it works as it ensures that the goroutine is removed from the WaitGroup when it finishes

	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Println("Part A", Num)

	mutex.Lock()     // Make sure that the go routines will wait at the barrier
	*count++         // Increment the counter to keep track of the number of goroutines that have reached the barrier
	if *count == 5 { // If all goroutines have reached the barrier
		close(barrier) // Closes the barrier channel to signal the goroutines to proceed, idea provided by chatGpt
	}
	mutex.Unlock() // Unlock the mutex to allow the goroutines to proceed

	// Wait for the signal to proceed
	<-barrier

	fmt.Println("Part B", Num)
}

func main() {
	threadCount := 5
	var wg sync.WaitGroup
	count := 0
	mutex := &sync.Mutex{}     // Mutex to protect shared counter
	barrier := make(chan bool) // Channel to act as the barrier for the go routines

	wg.Add(threadCount)
	for N := 0; N < threadCount; N++ {
		go WorkWithRendezvous(&wg, N, &count, barrier, mutex)
	}

	wg.Wait() // Wait for all goroutines to finish
}

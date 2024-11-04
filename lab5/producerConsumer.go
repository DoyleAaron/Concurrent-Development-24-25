package main

import (
	"fmt"
	"sync"
)

// Create a barrier data type
type barrier struct {
	theChan chan bool
	theLock sync.Mutex
	total   int
	count   int
}

// creates a properly initialised barrier
// N== number of threads (go Routines)
func createBarrier(N int) *barrier {
	theBarrier := &barrier{
		theChan: make(chan bool, 1), // Buffered channel of size 1 to make sure it is non-blocking
		total:   N,                  // Total number of items in the channel
		count:   0,                  // Current amount of items in the channel
	}
	return theBarrier
}

func producer(barrier *barrier, wg *sync.WaitGroup) {
	barrier.theLock.Lock()             // Locking the barrier to make sure that only one value can enter
	if barrier.count < barrier.total { // Checking if the channel is full
		barrier.count++
		fmt.Println("Producer added to channel, current count: ", barrier.count) // Printing to show it has been added
		barrier.theLock.Unlock()                                                 // Unlocking to allow the next value to enter
		barrier.theChan <- true                                                  // Signal that an item has been produced
	} else {
		barrier.theLock.Unlock()                            // Unlocking to allow the next value to enter
		fmt.Println("Channel is full, producer is waiting") // Print if the channel is full
	}
	wg.Done()
}

func consumer(barrier *barrier, wg *sync.WaitGroup) {
	<-barrier.theChan      // Wait for a signal that an item is available
	barrier.theLock.Lock() // Locking the barrier to make sure that only one value can enter
	if barrier.count > 0 { // Checking if the channel is empty
		barrier.count--                                                               // Decrementing the count to show that an item has been consumed
		fmt.Println("Consumer consumed from channel, current count: ", barrier.count) // Printing to show it has been consumed
	} else {
		fmt.Println("Channel is empty, consumer is waiting") // Print if the channel is empty
	}
	barrier.theLock.Unlock() // Unlocking to allow the next value to enter
	wg.Done()
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	barrier := createBarrier(5) // Create a barrier with a total of 5 items

	for i := 0; i < totalRoutines; i++ {
		go producer(barrier, &wg) // Start the producer goroutines
		go consumer(barrier, &wg) // Start the consumer goroutines
	}

	wg.Wait() // Wait for all goroutines to complete
}

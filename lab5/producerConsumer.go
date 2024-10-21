//producerConsumer.go code
// See readme for license

//--------------------------------------------
// Author: Aaron Doyle
// Created on 14/10/2024
// Modified by:
// Description:
// A simple example of producer consumer code using go
// Issues:
// figure out how to structure the producers and keep track of when the channel is full
//--------------------------------------------

package main

import (
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
func createBarrier(N int) barrier {
	theBarrier := barrier{
		theChan: make(chan bool),
		total:   N,
		count:   0,
	}
	return theBarrier
}

func producer() {

}

func consumer() {

}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	barrier := createBarrier(5)
	for i := 1; i < totalRoutines; i++ {

	}

} //end-main

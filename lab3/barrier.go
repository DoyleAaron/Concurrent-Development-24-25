// See readme for license

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Aaron Doyle
// Issues:
// The barrier is not implemented!
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, arrived *int, max int, wg *sync.WaitGroup, sharedLock *sync.Mutex, theChan chan bool) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	//we wait here until everyone has completed part A
	sharedLock.Lock()
	*arrived++
	if *arrived == 10 { //last to arrive -signal others to go
		sharedLock.Unlock() //unlock before any potentially blocking code
		theChan <- true
		<-theChan
	} else { //not all here yet we wait until signal
		sharedLock.Unlock() //unlock before any potentially blocking code
		<-theChan
		theChan <- true //once we get through send signal to next routine to continue
	} //end of if-else
	//we wait here until everyone has completed part A
	fmt.Println("Part B", goNum)
	wg.Done()
	return true
}

func main() {
	totalRoutines := 10
	arrived := 0
	var wg sync.WaitGroup
	max := 5
	wg.Add(totalRoutines)
	//we will need some of these
	var theLock sync.Mutex
	theChan := make(chan bool)
	theLock.Lock()
	count := 0
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &arrived, max, &wg, &theLock, theChan)
		if count > 9 {
			i = 0
		}
	}
	theLock.Unlock()

	wg.Wait() //wait for everyone to finish before exiting
}

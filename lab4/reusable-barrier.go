//barrier2.go Template Code
//See readme for license
//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Aaron Doyle, Ronan Green, Michael Cullen
// Description:
// A simple barrier implemented using mutex and unbuffered channel
// Issues:
// None I hope
//1. Change mutex to atomic variable
//2. Make it a reusable barrier
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, arrived *int, max int, wg *sync.WaitGroup, sharedLock *sync.Mutex, theChan chan bool, theChan2 chan bool) bool {
	for i := 1; i < 3; i++ {
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

		// everything is waiting here until the threads are finished
		sharedLock.Lock()
		*arrived--
		if *arrived == 0 { // checking if all have arrived
			sharedLock.Unlock() // unlocking to prevent a deadlock
			theChan2 <- true
			<-theChan2 // Setting it to wait for a signal
		} else {
			sharedLock.Unlock() // unlocking to prevent a deadlock
			<-theChan2
			theChan2 <- true // This is sending a signal to the next routine
		}
		fmt.Println("PartB", goNum)
	}
	wg.Done()
	return true
} //end-doStuff

func main() {
	totalRoutines := 10
	max := 5
	arrived := 0
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these
	var theLock sync.Mutex
	theChan := make(chan bool)     //use unbuffered channel in place of semaphore
	theChan2 := make(chan bool)    //creating a second channel to stop collision of data
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &arrived, max, &wg, &theLock, theChan, theChan2)
	}
	wg.Wait() //wait for everyone to finish before exiting
} //end-main

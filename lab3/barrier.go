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
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup, theLock *sync.Mutex, count *int, theSem *semaphore.Weighted, ctx context.Context) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	theLock.Lock()
	*count++
	if *count == max {
		theLock.Unlock()
		theChan <- true
		<-theChan
	} else {
		theLock.Unlock()
		<-theChan
		theChan <- true
	}
	//we wait here until everyone has completed part A
	theLock.Lock()
	*count--
	theLock.Unlock()
	fmt.Println("Part B", goNum)
	wg.Done()
	return true
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these
	ctx := context.TODO()
	var theLock sync.Mutex
	sem := semaphore.NewWeighted(int64(totalRoutines))
	theLock.Lock()
	sem.Acquire(ctx, 1)
	count := 0
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &wg, &theLock, &count, sem, ctx)
		if count > 9 {
			i = 0
		}
	}
	sem.Release(1)
	theLock.Unlock()

	wg.Wait() //wait for everyone to finish before exiting
}

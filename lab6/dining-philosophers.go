// // dining-philosophers.go code
// // See readme for license
// //--------------------------------------------
// // Author: Joseph Kehoe
// // Created on 5/11/2024
// // Modified by: Aaron Doyle
// // Description:
// // Solving the dining philosophers problem
// // Issues:
// // Figure out how all of the philosophers can eat their food without causing a deadlock when they all try to pick up the same fork at the same time
// //--------------------------------------------

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func think(index int) {
	var X time.Duration
	X = time.Duration(rand.Intn(2)) // Random time amount to simulate thinking
	time.Sleep(X * time.Second)     //wait random time amount
	fmt.Println("Phil: ", index, "was thinking")
}

func eat(index int) {
	var X time.Duration
	X = time.Duration(rand.Intn(2)) // Random time amount to simulate eating
	time.Sleep(X * time.Second)     //wait random time amount
	fmt.Println("Philosopher", index, "was eating")
}

func getForks(index int, forks map[int]chan bool) {
	if index%2 == 0 { // I'm splitting the even and odd philosophers to prevent deadlock so that they don't all try to pick up the same fork at the same time by splitting them into even and odd
		forks[index] <- true       // The left fork is the current philosopher's index and they pick this up first
		forks[(index+1)%5] <- true // We get the right fork by adding 1 to the current index and taking the remainder of 5 to ensure it's less than the number of philosophers
	} else { // Here I'm doing the same thing as above but in reverse for the odd philosophers so they grab the right fork first to prevent deadlock
		forks[(index+1)%5] <- true // The odd numbered philosophers pick up the right fork first
		forks[index] <- true       // Then they pick up the left fork to prevent deadlock
	}
}

func putForks(index int, forks map[int]chan bool) {
	if index%2 == 0 { // This follows the same logic as the getForks function above by splitting the even and odd philosophers to prevent deadlock
		<-forks[index]       // The evens put down the left fork first
		<-forks[(index+1)%5] // and then they put down the right fork this way it prevents deadlock as they don't all try to put down the same fork at the same time
	} else { // This is the same as the evens but in reverse for the odds to prevent deadlock
		<-forks[(index+1)%5] // The odds put down the right fork first so theyre the opposite of the evens which prevents deadlock
		<-forks[index]       // then they put down the left fork
	}
}

func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool) {
	think(index)           // Calling the think function to simulate the philosopher thinking
	getForks(index, forks) // Making the philosopher pick up the forks
	eat(index)             // Making the philosopher eat
	putForks(index, forks) // Making the philosopher put down the forks
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	philCount := 5    // Number of philosophers
	wg.Add(philCount) // Adding the philosophers to the wait group

	forks := make(map[int]chan bool) // Creating a map of forks
	for k := range philCount {
		forks[k] = make(chan bool, 1)
	} //set up forks

	for N := range philCount {
		go doPhilStuff(N, &wg, forks)
	} //start philosophers

	wg.Wait()
}

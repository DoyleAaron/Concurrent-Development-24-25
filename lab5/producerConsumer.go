//producerConsumer.go code
//Copyright (C) 2024 Mr. Aaron Doyle

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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

func producer() {

}

func consumer() {

}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)

	for i := 1; i < totalRoutines; i++ {

	}

} //end-main

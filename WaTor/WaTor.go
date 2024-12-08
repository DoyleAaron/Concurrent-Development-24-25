// Wa-Tor.go code
// See readme for license
//--------------------------------------------
// Author: Joseph Kehoe
// Created on 18/11/2024
// Modified by: Aaron Doyle
// Description:
// Creating a working version of the Wa-Tor simulation
// Issues:
// Figure out how to create a working version of the Wa-Tor simulation
//--------------------------------------------

package main

import (
	"WaTor/wator" // Import the wator package

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// Colours
	fishColour := rl.Green
	sharkColour := rl.Red
	waterColour := rl.Blue

	// Run the simulation
	wator.RunSimulation(fishColour, sharkColour, waterColour)
}

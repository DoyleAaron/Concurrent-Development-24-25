// // Wa-Tor.go code
// // See readme for license
// //--------------------------------------------
// // Author: Joseph Kehoe
// // Created on 18/11/2024
// // Modified by: Aaron Doyle
// // Description:
// // Creating a working version of the Wa-Tor simulation
// // Issues:
// // Figure out how to create a working version of the Wa-Tor simulation
// //--------------------------------------------

// Package main implements a Wa-Tor simulation using Raylib in the Go programming language
package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// DrawFish function
// Parameters: x, y, width, height int, fishColour rl.Color
// Returns: None
// Description: Draws a fish at the given x and y coordinates with the given width and height and color
func DrawFish(x, y, width, height int, fishColour rl.Color) {
	rl.DrawRectangle(int32(x), int32(y), int32(width), int32(height), fishColour)
}

// DrawShark function
// Parameters: x, y, width, height int, sharkColour rl.Color
// Returns: None
// Description: Draws a shark at the given x and y coordinates with the given width and height and color
func DrawShark(x, y, width, height int, sharkColour rl.Color) {
	rl.DrawRectangle(int32(x), int32(y), int32(width), int32(height), sharkColour)
}

// DrawWater function
// Parameters: x, y, width, height int, waterColour rl.Color
// Returns: None
// Description: Draws water at the given x and y coordinates with the given width and height and color
func DrawWater(x, y, width, height int, waterColour rl.Color) {
	rl.DrawRectangle(int32(x), int32(y), int32(width), int32(height), waterColour)
}

// InitialPositions function
// Parameters: xdim, ydim, NumShark, NumFish int
// Returns: grid []int
// Description: Initializes the positions of the fish and sharks in the grid so that they are randomly placed around the grid
func InitialPositions(xdim, ydim, NumShark, NumFish int) []int {

	// ChatGPT helped me with the grid part of the code as I was confused as to how to initially set the fish and shark positions
	// It works by creating a grid with the dimensions of xdim and ydim and the number of sharks and fish, and then shuffling the grid to simulate randomly placing the fish and sharks around the screen
	grid := make([]int, xdim*ydim)
	for i := 0; i < NumShark; i++ {
		grid[i] = 1 // 1 represents sharks
	}
	for i := NumShark; i < NumShark+NumFish; i++ {
		grid[i] = 2 // 2 represents fish
	}

	// This mixes up the order of the grid to simulate randomly placing the fish and sharks around the screen
	rand.Shuffle(len(grid), func(i, j int) { grid[i], grid[j] = grid[j], grid[i] })

	return grid
}

// UpdatePositions function
// Parameters: grid []int, xdim, ydim int (current grid)
// Returns: newGrid []int (updated grid)
// Description: Updates the positions of the fish and sharks based off of the rules in the specification
func UpdatePositions(grid []int, xdim, ydim int) []int {
	// Create a new grid to store the updated positions of the fish and sharks
	newGrid := make([]int, xdim*ydim)

	for i := 0; i < xdim; i++ {
		for j := 0; j < ydim; j++ {
			// Get the current cell value
			cellValue := grid[j*xdim+i] // ChatGPT gave me this calculation to get the current cell value in a 2D array

			// Check if the cell value is a shark
			if cellValue == 1 {

			}
		}
	}

	// Return the updated grid
	return newGrid

}

// Main Class
// Description: Main function that handles the variables and calls the functions for the simulation
func main() {
	// Constants
	xdim := 100
	ydim := 100
	windowXSize := 800
	windowYSize := 600
	cellXSize := windowXSize / xdim
	cellYSize := windowYSize / ydim
	NumShark := 20
	NumFish := 100

	// Colors
	fishColour := rl.Green
	sharkColour := rl.Red
	waterColour := rl.Blue

	// Initialize the window
	rl.InitWindow(int32(windowXSize), int32(windowYSize), "Raylib Wa-Tor world")
	defer rl.CloseWindow() // Ensure the window is closed on exit

	// Initialize grid
	grid := InitialPositions(xdim, ydim, NumShark, NumFish)

	// Simulation loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		// Clear the screen
		rl.ClearBackground(rl.RayWhite)

		// Draw the grid
		for i := 0; i < xdim; i++ {
			for j := 0; j < ydim; j++ {
				x := i * cellXSize
				y := j * cellYSize
				cellValue := grid[j*xdim+i]

				if cellValue == 1 {
					DrawShark(x, y, cellXSize, cellYSize, sharkColour)
				} else if cellValue == 2 {
					DrawFish(x, y, cellXSize, cellYSize, fishColour)
				} else {
					DrawWater(x, y, cellXSize, cellYSize, waterColour)
				}
			}
		}

		// Update the grid
		//grid = UpdatePositions(grid, xdim, ydim)

		rl.EndDrawing()
	}
}

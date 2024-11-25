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

// Main Class
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

	// Colours of the fish, shark, and water
	fishColour := rl.Green
	sharkColour := rl.Red
	waterColour := rl.Blue

	// Random seed for entity placement
	rand.Int()

	// ChatGPT helped me with the grid part of the code as I was confused as to how to initially set the fish and shark positions
	// It works by creating a grid with the dimensions of xdim and ydim and the number of sharks and fish, and then shuffling the grid to simulate randomly placing the fish and sharks around the screen
	grid := make([]int, xdim*ydim)
	for i := 0; i < NumShark; i++ {
		grid[i] = 1
	}
	for i := NumShark; i < NumShark+NumFish; i++ {
		grid[i] = 2
	}

	rand.Shuffle(len(grid), func(i, j int) { grid[i], grid[j] = grid[j], grid[i] })

	// Initialize the window
	rl.InitWindow(int32(windowXSize), int32(windowYSize), "Raylib Wa-Tor world")
	defer rl.CloseWindow() // Close window when done

	// Logic loop
	// This loop will handle the running of the simulation and will run until the window is closed
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		// These for loops will iterate through the grid and draw the fish, sharks, and water to the screen
		for i := 0; i < xdim; i++ {
			for k := 0; k < ydim; k++ {
				x := i * cellXSize
				y := k * cellYSize
				cellValue := grid[k*xdim+i]
				if cellValue == 1 {
					DrawShark(x, y, cellXSize, cellYSize, sharkColour)
				} else if cellValue == 2 {
					DrawFish(x, y, cellXSize, cellYSize, fishColour)
				} else {
					DrawWater(x, y, cellXSize, cellYSize, waterColour)
				}
			}
		}

		rl.EndDrawing()
	}
}

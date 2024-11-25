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

	// ChatGPT helped m
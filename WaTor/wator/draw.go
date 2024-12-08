// draw.go
// Package wator contains drawing logic for the simulation.

package wator

import rl "github.com/gen2brain/raylib-go/raylib"

// DrawFish function
// Parameters: x, y, width, height int, fishColour rl.Color
// Returns: None
// Description: Draws a fish at the given x and y coordinates with the given width and height and colour
func DrawFish(x, y, width, height int, fishColour rl.Color) {
	rl.DrawRectangle(int32(x), int32(y), int32(width), int32(height), fishColour)
}

// DrawShark function
// Parameters: x, y, width, height int, sharkColour rl.Color
// Returns: None
// Description: Draws a shark at the given x and y coordinates with the given width and height and colour
func DrawShark(x, y, width, height int, sharkColour rl.Color) {
	rl.DrawRectangle(int32(x), int32(y), int32(width), int32(height), sharkColour)
}

// DrawWater function
// Parameters: x, y, width, height int, waterColour rl.Color
// Returns: None
// Description: Draws water at the given x and y coordinates with the given width and height and colour
func DrawWater(x, y, width, height int, waterColour rl.Color) {
	rl.DrawRectangle(int32(x), int32(y), int32(width), int32(height), waterColour)
}

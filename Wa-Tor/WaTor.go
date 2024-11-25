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

package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// Initialize variables
	xdim := 100
	ydim := 100
	windowXSize := 800
	windowYSize := 600
	cellXSize := windowXSize / xdim
	cellYSize := windowYSize / ydim
	// NumShark := 100
	// NumFish := 100
	// FishBreed := 3
	// SharkBreed := 3
	// Starve := 3

	// Initialize the window
	rl.InitWindow(int32(windowXSize), int32(windowYSize), "Raylib Wa-Tor world")
	defer rl.CloseWindow() // Close window when done

	// Loop to draw the rectangles that keeps going until the escape key is pressed
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// Draw rectangles (each representing a fish, shark, or empty space)
		for i := 0; i < xdim; i++ {
			for k := 0; k < ydim; k++ {
				recX := float32(i * cellXSize)
				recY := float32(k * cellYSize)
				recWidth := float32(80)  // Width of the rectangle
				recHeight := float32(60) // Height of the rectangle
				//id := i*1 - k

				color := rl.Green
				if rand.Intn(2) == 0 {
					color = rl.Blue
				}

				// Draw the rectangle
				rl.DrawRectangle(int32(recX), int32(recY), int32(recWidth), int32(recHeight), color)
			}
		}

		// End drawing and display the window
		rl.EndDrawing()
	}
}

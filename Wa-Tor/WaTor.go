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
				fishX := float32(i * cellXSize)  // X position of the fish
				fishY := float32(k * cellYSize)  // Y position of the fish
				sharkX := float32(i * cellXSize) // X position of the shark
				sharkY := float32(k * cellYSize) // Y position of the shark
				Width := float32(80)             // Width of the rectangle
				Height := float32(60)            // Height of the rectangle
				fishColour := rl.Green
				sharkColour := rl.Blue

				if(){
					
				}
		}

		// End drawing and display the window
		rl.EndDrawing()
	}
}

func drawFish(fishX, fishY, Width, Height float32, fishColour rl.Color) {
	rl.DrawRectangle(int32(fishX), int32(fishY), int32(Width), int32(Height), fishColour)
}

func drawShark(sharkX, sharkY, Width, Height float32, sharkColour rl.Color) {
	rl.DrawRectangle(int32(sharkX), int32(sharkY), int32(Width), int32(Height), sharkColour)
}

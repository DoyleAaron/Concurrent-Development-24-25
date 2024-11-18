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

	// Initialize the window
	rl.InitWindow(int32(windowXSize), int32(windowYSize), "Raylib Wa-Tor world")
	defer rl.CloseWindow() // Close window when done

	// Loop to draw the rectangles
	for !rl.WindowShouldClose() {
		// Start drawing
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// Draw rectangles (each representing a fish, shark, or empty space)
		for i := 0; i < xdim; i++ {
			for k := 0; k < ydim; k++ {
				recX := float32(i * cellXSize)
				recY := float32(k * cellYSize)
				recWidth := float32(80)  // Width of the rectangle
				recHeight := float32(60) // Height of the rectangle
				id := i*1 - k

				// Set color based on the ID
				var color rl.Color
				if id%2 == 0 {
					color = rl.Green // Fish
				} else {
					color = rl.Blue // Empty space
				}

				// Draw the rectangle
				rl.DrawRectangle(int32(recX), int32(recY), int32(recWidth), int32(recHeight), color)
			}
		}

		// End drawing and display the window
		rl.EndDrawing()
	}
}

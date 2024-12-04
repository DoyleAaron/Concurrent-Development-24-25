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

// Constant variables
const (
	xdim        = 100                // Width of the grid
	ydim        = 100                // Height of the grid
	windowXSize = 800                // Width of the window
	windowYSize = 600                // Height of the window
	cellXSize   = windowXSize / xdim // Width of each cell
	cellYSize   = windowYSize / ydim // Height of each cell
	NumShark    = 500                // The number of sharks in the simulation
	NumFish     = 100                // The number of fish in the simulation
	Starve      = 900                // The number of turns it takes for a shark to starve
	SharkBreed  = 114                // The number of turns it takes for a shark to breed
	FishBreed   = 160                // The number of turns it takes for a fish to breed
)

// Cell struct
// Parameters: Type, BreedTime, StarveTime
// Returns: None
// Description: This is to give the cells in the grid a type, breed time and starve time
type Cell struct {
	Type             int // 0 = water, 1 = shark, 2 = fish
	BreedTime        int // The number of turns it takes for the fish or shark to breed
	StarveTime       int // The number of turns it takes for the shark to starve
	CurrentBreedTime int // The current number of turns the fish or shark has been alive
}

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
// Returns: grid [][]Cell
// Description: Initializes the positions of the fish and sharks in the grid so that they are randomly placed around the grid
func InitialPositions(xdim, ydim, NumShark int, NumFish int, Starve int, FishBreed int, SharkBreed int, rnd *rand.Rand) [][]Cell {

	// ChatGPT helped me with the grid part of the code as I was confused as to how to initially set the fish and shark positions
	// It works by creating a grid with the dimensions of xdim and ydim and the number of sharks and fish, then it creates a 1D slice to represent the flattened grid for random placement.
	// Once the fish and sharks are placed, it maps the flat grid back into the 2D grid so it can be displayed
	grid := make([][]Cell, ydim)
	for i := 0; i < ydim; i++ {
		grid[i] = make([]Cell, xdim)
	}

	// Create a 1D slice to represent the flattened grid for random placement
	flatGrid := make([]Cell, xdim*ydim)

	for i := 0; i < NumShark; i++ {
		flatGrid[i] = Cell{Type: 1, BreedTime: SharkBreed, StarveTime: Starve, CurrentBreedTime: 0} // 1 represents sharks
	}

	for i := NumShark; i < NumShark+NumFish; i++ {
		flatGrid[i] = Cell{Type: 2, BreedTime: FishBreed, CurrentBreedTime: 0} // 2 represents fish
	}

	// This mixes up the order of the grid to simulate randomly placing the fish and sharks around the screen
	rnd.Shuffle(len(flatGrid), func(i, j int) { flatGrid[i], flatGrid[j] = flatGrid[j], flatGrid[i] })

	// Map the flat grid back into the 2D grid
	for i := 0; i < ydim; i++ {
		for j := 0; j < xdim; j++ {
			grid[i][j] = flatGrid[i*xdim+j]
		}
	}

	return grid
}

// UpdatePositions function
// Parameters: grid [][]Cell, xdim, ydim int (current grid)
// Returns: newGrid [][]Cell (updated grid)
// Description: Updates the positions of the fish and sharks based off of the rules in the specification
func UpdatePositions(grid [][]Cell, xdim, ydim int, rnd *rand.Rand) [][]Cell {
	// Create a new grid to store the updated positions of the fish and sharks
	newGrid := make([][]Cell, ydim)
	for i := 0; i < ydim; i++ {
		newGrid[i] = make([]Cell, xdim)
		copy(newGrid[i], grid[i])
	}

	for i := 0; i < xdim; i++ {
		for j := 0; j < ydim; j++ {
			currentCell := grid[j][i] // Get the current cell and its information

			// Check if the cell contains a shark
			if currentCell.Type == 1 {
				// Check if the shark starves
				if currentCell.StarveTime <= 0 {
					newGrid[j][i] = Cell{Type: 0} // The shark has starved and is removed from the grid
					continue                      // Skip further processing for this shark
				}

				moved := false

				// Check if there is a fish to the north
				if j > 0 && grid[j-1][i].Type == 2 && !moved {
					newGrid[j-1][i] = Cell{
						Type:             1,
						BreedTime:        currentCell.BreedTime,
						StarveTime:       Starve, // Reset starvation timer
						CurrentBreedTime: currentCell.CurrentBreedTime + 1,
					}
					newGrid[j][i] = Cell{Type: 0} // Turn current position into water
					moved = true
				} else if i < xdim-1 && grid[j][i+1].Type == 2 && !moved { // Check if there is a fish to the east
					newGrid[j][i+1] = Cell{
						Type:             1,
						BreedTime:        currentCell.BreedTime,
						StarveTime:       Starve,
						CurrentBreedTime: currentCell.CurrentBreedTime + 1,
					}
					newGrid[j][i] = Cell{Type: 0} // Turn current position into water
					moved = true
				} else if j < ydim-1 && grid[j+1][i].Type == 2 && !moved { // Check if there is a fish to the south
					newGrid[j+1][i] = Cell{
						Type:             1,
						BreedTime:        currentCell.BreedTime,
						StarveTime:       Starve,
						CurrentBreedTime: currentCell.CurrentBreedTime + 1,
					}
					newGrid[j][i] = Cell{Type: 0} // Turn current position into water
					moved = true
				} else if i > 0 && grid[j][i-1].Type == 2 && !moved { // Check if there is a fish to the west
					newGrid[j][i-1] = Cell{
						Type:             1,
						BreedTime:        currentCell.BreedTime,
						StarveTime:       Starve,
						CurrentBreedTime: currentCell.CurrentBreedTime + 1,
					}
					newGrid[j][i] = Cell{Type: 0}
					moved = true
				}

				// If the shark hasn't moved, decrement its StarveTime and handle breeding/moving
				if !moved {
					newGrid[j][i].StarveTime--
					freeSpace := []struct{ x, y int }{}

					// Check for empty cells around the shark and if there are any add them to the freeSpace slice
					if j > 0 && grid[j-1][i].Type == 0 {
						freeSpace = append(freeSpace, struct{ x, y int }{i, j - 1})
					}
					if i < xdim-1 && grid[j][i+1].Type == 0 {
						freeSpace = append(freeSpace, struct{ x, y int }{i + 1, j})
					}
					if j < ydim-1 && grid[j+1][i].Type == 0 {
						freeSpace = append(freeSpace, struct{ x, y int }{i, j + 1})
					}
					if i > 0 && grid[j][i-1].Type == 0 {
						freeSpace = append(freeSpace, struct{ x, y int }{i - 1, j})
					}

					//  If there are empty cells around the shark, move to one of them
					if len(freeSpace) > 0 {
						randDirection := rnd.Intn(len(freeSpace)) // I'm using this to randomly choose a direction for the shark to move
						chosenDirection := freeSpace[randDirection]

						// Handle breeding or moving
						if currentCell.CurrentBreedTime == currentCell.BreedTime {
							newGrid[chosenDirection.y][chosenDirection.x] = Cell{
								Type:             1,
								BreedTime:        SharkBreed,
								StarveTime:       Starve,
								CurrentBreedTime: 0, // Reset breed time
							}
							newGrid[j][i] = Cell{
								Type:             1,
								BreedTime:        currentCell.BreedTime,
								StarveTime:       currentCell.StarveTime - 1,
								CurrentBreedTime: 0,
							}
						} else {
							newGrid[chosenDirection.y][chosenDirection.x] = Cell{
								Type:             1,
								BreedTime:        currentCell.BreedTime,
								StarveTime:       currentCell.StarveTime - 1,
								CurrentBreedTime: currentCell.CurrentBreedTime + 1,
							}
							newGrid[j][i] = Cell{Type: 0}
						}
					} else {
						// Decrement StarveTime if the shark hasn't moved
						newGrid[j][i].StarveTime--
						newGrid[j][i].CurrentBreedTime++
					}
				}
			} else if currentCell.Type == 2 { // Check if the cell contains a fish
				freeSpace := []struct{ x, y int }{}

				// Check for empty cells around the fish and if there are any add them to the freeSpace slice
				if j > 0 && grid[j-1][i].Type == 0 {
					freeSpace = append(freeSpace, struct{ x, y int }{i, j - 1})
				}
				if i < xdim-1 && grid[j][i+1].Type == 0 {
					freeSpace = append(freeSpace, struct{ x, y int }{i + 1, j})
				}
				if j < ydim-1 && grid[j+1][i].Type == 0 {
					freeSpace = append(freeSpace, struct{ x, y int }{i, j + 1})
				}
				if i > 0 && grid[j][i-1].Type == 0 {
					freeSpace = append(freeSpace, struct{ x, y int }{i - 1, j})
				}

				if len(freeSpace) > 0 {
					randDirection := rnd.Intn(len(freeSpace))
					chosenDirection := freeSpace[randDirection]

					// This is to check if the fish can breed
					if currentCell.CurrentBreedTime == currentCell.BreedTime {
						newGrid[chosenDirection.y][chosenDirection.x] = Cell{
							Type:             2,
							BreedTime:        currentCell.BreedTime,
							CurrentBreedTime: 0, // Reset breed time
						}
						newGrid[j][i] = Cell{
							Type:             2,
							BreedTime:        currentCell.BreedTime,
							CurrentBreedTime: 0,
						}
						continue
					} else {
						newGrid[chosenDirection.y][chosenDirection.x] = Cell{
							Type:             2,
							BreedTime:        currentCell.BreedTime,
							CurrentBreedTime: currentCell.CurrentBreedTime + 1,
						}
						newGrid[j][i] = Cell{Type: 0}
					}
				} else {
					continue // If the fish can't move then continue to the next cell
				}
			} else {
				continue // If the cell is water then continue to the next cell
			}
		}
	}
	return newGrid
}

// Main Class
// Description: Main function that handles the variables and calls the functions for the simulation
func main() {

	// Colors
	fishColour := rl.Green
	sharkColour := rl.Red
	waterColour := rl.Blue

	rnd := rand.New(rand.NewSource(42))

	// Initialize the window
	rl.InitWindow(int32(windowXSize), int32(windowYSize), "Raylib Wa-Tor Simulation")
	defer rl.CloseWindow() // Ensure the window is closed on exit

	// Initialize grid
	grid := InitialPositions(xdim, ydim, NumShark, NumFish, Starve, FishBreed, SharkBreed, rnd)

	// Simulation loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		// Draw the grid
		for i := 0; i < xdim; i++ {
			for j := 0; j < ydim; j++ {
				x := i * cellXSize
				y := j * cellYSize
				cell := grid[j][i]

				if cell.Type == 1 { // Draw a shark
					DrawShark(x, y, cellXSize, cellYSize, sharkColour)
				} else if cell.Type == 2 { // Draw a fish
					DrawFish(x, y, cellXSize, cellYSize, fishColour)
				} else { // Draw water
					DrawWater(x, y, cellXSize, cellYSize, waterColour)
				}
			}
		}

		// Update the grid
		grid = UpdatePositions(grid, xdim, ydim, rnd)

		rl.EndDrawing()
	}
}

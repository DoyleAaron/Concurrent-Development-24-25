// simulation.go
// Package wator contains the main simulation logic.

package wator

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// UpdatePositions function
// Parameters: grid [][]Cell, xdim, ydim int (current grid)
// Returns: newGrid [][]Cell (updated grid)
// Description: Updates the positions of the fish and sharks based off of the rules in the specification
func UpdatePositions(grid [][]Cell, xdim, ydim int, rnd *rand.Rand) [][]Cell {

	// Reset visited flags for all cells
	for i := 0; i < ydim; i++ {
		for j := 0; j < xdim; j++ {
			grid[i][j].Visited = 0
		}
	}

	for i := 0; i < xdim; i++ {
		for j := 0; j < ydim; j++ {
			currentCell := grid[j][i] // Get the current cell and its information
			if currentCell.Visited == 1 {
				continue // Skip further processing for this cell
			}

			// SHARK LOGIC COMMENCES HERE
			// Overview:
			// Check if the cell contains a shark
			// Then check if there is a fish to the north, east, south or west of the shark
			// If there is a fish in any of these directions, move the shark to that cell and turn the current cell into water
			// If there is no fish in any of these directions, decrement the shark's StarveTime and check for empty cells around the shark
			// If there are empty cells around the shark, move the shark to one of them randomly and leave water in the current cell
			// If the shark can breed, move the shark into the empty cell and leave the current cell as a new shark
			if currentCell.Type == 1 {
				// Check if the shark starves
				if currentCell.StarveTime <= 0 {
					grid[j][i] = Cell{Type: 0, Visited: 1} // The shark has starved and is removed from the grid
					continue                               // Skip further processing for this shark
				}

				moved := false
				fishDirections := []struct{ x, y int }{}

				// Check if there is a fish to the north, east, south or west of the shark
				if j > 0 && grid[j-1][i].Type == 2 {
					fishDirections = append(fishDirections, struct{ x, y int }{j - 1, i})
				}
				if i < xdim-1 && grid[j][i+1].Type == 2 {
					fishDirections = append(fishDirections, struct{ x, y int }{j, i + 1})
				}
				if j < ydim-1 && grid[j+1][i].Type == 2 {
					fishDirections = append(fishDirections, struct{ x, y int }{j + 1, i})
				}
				if i > 0 && grid[j][i-1].Type == 2 {
					fishDirections = append(fishDirections, struct{ x, y int }{j, i - 1})
				}

				if len(fishDirections) > 0 {
					randDirection := rnd.Intn(len(fishDirections))
					chosenDirection := fishDirections[randDirection]

					grid[chosenDirection.x][chosenDirection.y] = Cell{
						Type:             1,
						BreedTime:        currentCell.BreedTime,
						StarveTime:       Starve, // Reset starvation timer
						CurrentBreedTime: currentCell.CurrentBreedTime + 1,
						Visited:          1,
					}
					grid[j][i] = Cell{Type: 0, Visited: 1} // Turn current position into water
					moved = true
				}

				// If the shark hasn't moved, decrement its StarveTime and handle breeding/moving
				if !moved {
					grid[j][i].StarveTime--
					freeSpace := []struct{ x, y int }{}

					// Check for empty cells around the shark and if there are any add them to the freeSpace slice
					if j > 0 && grid[j-1][i].Type == 0 {
						freeSpace = append(freeSpace, struct{ x, y int }{j - 1, i})
					}
					if i < xdim-1 && grid[j][i+1].Type == 0 {
						freeSpace = append(freeSpace, struct{ x, y int }{j, i + 1})
					}
					if j < ydim-1 && grid[j+1][i].Type == 0 {
						freeSpace = append(freeSpace, struct{ x, y int }{j + 1, i})
					}
					if i > 0 && grid[j][i-1].Type == 0 {
						freeSpace = append(freeSpace, struct{ x, y int }{j, i - 1})
					}

					//  If there are empty cells around the shark, move to one of them
					if len(freeSpace) > 0 {
						randDirection := rnd.Intn(len(freeSpace)) // I'm using this to randomly choose a direction for the shark to move
						chosenDirection := freeSpace[randDirection]

						// Handle breeding or moving
						if currentCell.CurrentBreedTime >= currentCell.BreedTime {
							grid[chosenDirection.y][chosenDirection.x] = Cell{
								Type:             1,
								BreedTime:        SharkBreed,
								StarveTime:       Starve,
								CurrentBreedTime: 0, // Reset breed time
								Visited:          1,
							}
							grid[j][i] = Cell{
								Type:             1,
								BreedTime:        currentCell.BreedTime,
								StarveTime:       currentCell.StarveTime - 1,
								CurrentBreedTime: 0,
								Visited:          1,
							}
						} else {
							grid[chosenDirection.y][chosenDirection.x] = Cell{
								Type:             1,
								BreedTime:        currentCell.BreedTime,
								StarveTime:       currentCell.StarveTime - 1,
								CurrentBreedTime: currentCell.CurrentBreedTime + 1,
								Visited:          1,
							}
							grid[j][i] = Cell{Type: 0, Visited: 1}
						}
					} else {
						// Decrement StarveTime if the shark hasn't moved
						grid[j][i].StarveTime--
						grid[j][i].CurrentBreedTime++
					}
				}
			}

			// FISH LOGIC COMMENCES HERE
			// Overview:
			// Check if the cell contains a fish
			// Then check if there are empty cells around the fish
			// If there are empty cells around the fish, move the fish to one of them and leave water in the current cell
			// If the fish can breed, move the fish into the empty cell and leave the current cell as a new fish
			if currentCell.Type == 2 { // Check if the cell contains a fish
				freeSpace := []struct{ x, y int }{}

				// Check for empty cells around the fish and if there are any add them to the freeSpace slice
				if j > 0 && grid[j-1][i].Type == 0 && grid[j-1][i].Visited == 0 {
					freeSpace = append(freeSpace, struct{ x, y int }{j - 1, i})
				}
				if i < xdim-1 && grid[j][i+1].Type == 0 && grid[j][i+1].Visited == 0 {
					freeSpace = append(freeSpace, struct{ x, y int }{j, i + 1})
				}
				if j < ydim-1 && grid[j+1][i].Type == 0 && grid[j+1][i].Visited == 0 {
					freeSpace = append(freeSpace, struct{ x, y int }{j + 1, i})
				}
				if i > 0 && grid[j][i-1].Type == 0 && grid[j][i-1].Visited == 0 {
					freeSpace = append(freeSpace, struct{ x, y int }{j, i - 1})
				}

				currentCell.CurrentBreedTime++ // Increment the current breed time for the fish

				// If there are empty cells around the fish, move to one of them
				if len(freeSpace) > 0 {
					randDirection := rnd.Intn(len(freeSpace))
					chosenDirection := freeSpace[randDirection]

					if currentCell.CurrentBreedTime >= currentCell.BreedTime {
						grid[chosenDirection.x][chosenDirection.y] = Cell{
							Type:             2,
							BreedTime:        currentCell.BreedTime,
							CurrentBreedTime: 0, // Reset breed time
							Visited:          1,
						}
						grid[j][i] = Cell{
							Type:             2,
							BreedTime:        currentCell.BreedTime,
							CurrentBreedTime: 0,
							Visited:          1,
						}
					} else {
						grid[j][i] = Cell{Type: 0, Visited: 1} // Turn current position into water
						grid[chosenDirection.x][chosenDirection.y] = Cell{
							Type:             2,
							BreedTime:        currentCell.BreedTime,
							CurrentBreedTime: currentCell.CurrentBreedTime + 1,
							Visited:          1,
						}
					}
				} else {
					grid[j][i].CurrentBreedTime = currentCell.CurrentBreedTime // Increment the current breed time for the fish
				}
			}
		}
	}

	return grid
}

// RunSimulation initializes and runs the Wa-Tor simulation.
func RunSimulation(fishColour, sharkColour, waterColour rl.Color) {
	rnd := rand.New(rand.NewSource(42))

	rl.InitWindow(int32(WindowXSize), int32(WindowYSize), "Raylib Wa-Tor Simulation")
	defer rl.CloseWindow()

	grid := InitialPositions(rnd)

	// Open CSV file for writing
	file, err := os.Create("./fps_log.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Timestamp", "FPS"})

	lastLogTime := time.Now()
	index := 0

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		// Draw the grid
		for i := 0; i < XDim; i++ {
			for j := 0; j < YDim; j++ {
				x := i * CellXSize
				y := j * CellYSize
				cell := grid[j][i]

				if cell.Type == 1 {
					DrawShark(x, y, CellXSize, CellYSize, sharkColour)
				} else if cell.Type == 2 {
					DrawFish(x, y, CellXSize, CellYSize, fishColour)
				} else {
					DrawWater(x, y, CellXSize, CellYSize, waterColour)
				}
			}
		}

		grid = UpdatePositions(grid, XDim, YDim, rnd)

		if time.Since(lastLogTime).Seconds() >= 1 {
			currentFPS := rl.GetFPS()
			index++
			writer.Write([]string{fmt.Sprintf("%d", index), fmt.Sprintf("%d", currentFPS)})
			writer.Flush()
			lastLogTime = time.Now()
		}

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
		time.Sleep(5 * time.Millisecond)
	}
}

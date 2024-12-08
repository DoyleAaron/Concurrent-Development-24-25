// grid.go
// Package wator contains grid initialization logic.

package wator

import "math/rand"

// InitialPositions function
// Parameters: xdim, ydim, NumShark, NumFish int, rnd *rand.Rand
// Returns: grid [][]Cell
// Description: Initializes the positions of the fish and sharks in the grid so that they are randomly placed around the grid
func InitialPositions(rnd *rand.Rand) [][]Cell {
	grid := make([][]Cell, YDim)
	for i := 0; i < YDim; i++ {
		grid[i] = make([]Cell, XDim)
	}

	// Create a 1D slice to represent the flattened grid for random placement
	flatGrid := make([]Cell, XDim*YDim)

	for i := 0; i < NumShark; i++ {
		flatGrid[i] = Cell{Type: 1, BreedTime: SharkBreed, StarveTime: Starve, CurrentBreedTime: 0} // 1 represents sharks
	}

	for i := NumShark; i < NumShark+NumFish; i++ {
		flatGrid[i] = Cell{Type: 2, BreedTime: FishBreed, CurrentBreedTime: 0} // 2 represents fish
	}

	// This mixes up the order of the grid to simulate randomly placing the fish and sharks around the screen
	rnd.Shuffle(len(flatGrid), func(i, j int) { flatGrid[i], flatGrid[j] = flatGrid[j], flatGrid[i] })

	// Map the flat grid back into the 2D grid
	for i := 0; i < YDim; i++ {
		for j := 0; j < XDim; j++ {
			grid[i][j] = flatGrid[i*XDim+j]
		}
	}
	return grid
}

// types.go
// Package wator contains types and constants for the Wa-Tor simulation.

package wator

// Constant variables
const (
	XDim        = 100                // Width of the grid
	YDim        = 100                // Height of the grid
	WindowXSize = 800                // Width of the window
	WindowYSize = 600                // Height of the window
	CellXSize   = WindowXSize / XDim // Width of each cell
	CellYSize   = WindowYSize / YDim // Height of each cell
	NumShark    = 300                // The number of sharks in the simulation
	NumFish     = 1000               // The number of fish in the simulation
	Starve      = 3                  // The number of turns it takes for a shark to starve
	SharkBreed  = 3                  // The number of turns it takes for a shark to breed
	FishBreed   = 10                 // The number of turns it takes for a fish to breed
)

// Cell represents a single grid cell in the Wa-Tor simulation.
// It holds the type of entity (water, shark, or fish) and their properties.
type Cell struct {
	Type             int // 0 = water, 1 = shark, 2 = fish
	BreedTime        int // The number of turns it takes for the fish or shark to breed
	StarveTime       int // The number of turns it takes for the shark to starve
	CurrentBreedTime int // The current number of turns the fish or shark has been alive
	Visited          int // This is to check if the cell has been visited and to ignore it if it has
}

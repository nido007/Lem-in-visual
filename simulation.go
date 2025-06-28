package main

import (
	"fmt"
	"strings"
)

// Ant represents a single ant in the simulation
type Ant struct {
	ID   int     // Unique number for this ant
	Path []*Room // The route this ant will follow
	Pos  int     // Current position along the path (0 = start)
}

// RunSimulation moves all ants from start to end, one turn at a time
func RunSimulation(farm *Farm, paths [][]*Room) {
	totalAnts := farm.AntCount
	numPaths := len(paths)

	if numPaths == 0 {
		return // No paths available
	}

	// Create queues of ants for each path
	queues := make([][]*Ant, numPaths)

	// Distribute ants among paths
	for i, path := range paths {
		// Calculate how many ants go on this path
		antsForThisPath := totalAnts / numPaths
		if i < totalAnts%numPaths {
			antsForThisPath++ // Give extra ant to first few paths
		}

		// Create ants for this path
		queues[i] = make([]*Ant, antsForThisPath)
		for j := 0; j < antsForThisPath; j++ {
			queues[i][j] = &Ant{
				ID:   0,    // Will be assigned when launched
				Path: path, // Route to follow
				Pos:  0,    // Start at beginning
			}
		}
	}

	nextID := 1                   // Next ant number to assign
	activeAnts := make([]*Ant, 0) // Ants currently moving
	finished := 0                 // How many ants have reached the end

	// Clear all room occupancy
	for _, room := range farm.Rooms {
		room.Occupied = false
	}

	// Main simulation loop - continue until all ants reach the end
	for finished < totalAnts {
		var moves []string // Moves made this turn

		// Phase 1: Launch new ants (one per path if possible)
		for i := range queues {
			if len(queues[i]) > 0 {
				// Take the next ant from this path's queue
				ant := queues[i][0]
				queues[i] = queues[i][1:] // Remove from queue

				ant.ID = nextID
				nextID++
				activeAnts = append(activeAnts, ant)
			}
		}

		// Phase 2: Move existing ants
		usedTunnels := make(map[string]bool) // Track which tunnels are used this turn

		for _, ant := range activeAnts {
			// Skip ants that have already reached the end
			if ant.Pos >= len(ant.Path)-1 {
				continue
			}

			// Determine current and next rooms
			var currentRoom *Room
			if ant.Pos == 0 {
				currentRoom = farm.Start
			} else {
				currentRoom = ant.Path[ant.Pos]
			}
			nextRoom := ant.Path[ant.Pos+1]

			// Create tunnel identifier
			tunnelID := currentRoom.Name + "->" + nextRoom.Name

			// Check if tunnel is already used this turn
			if usedTunnels[tunnelID] {
				continue // Can't use same tunnel twice in one turn
			}

			// Check if next room is occupied (except start/end)
			if nextRoom != farm.Start && nextRoom != farm.End && nextRoom.Occupied {
				continue // Room is occupied
			}

			// Free the previous room (if not start/end)
			if ant.Pos > 0 {
				prevRoom := ant.Path[ant.Pos]
				if prevRoom != farm.Start && prevRoom != farm.End {
					prevRoom.Occupied = false
				}
			}

			// Move the ant
			ant.Pos++
			usedTunnels[tunnelID] = true

			// Occupy the new room (if not start/end)
			if nextRoom != farm.Start && nextRoom != farm.End {
				nextRoom.Occupied = true
			}

			// Record the move
			moves = append(moves, PrintAntMove(ant.ID, nextRoom.Name))

			// Check if ant reached the end
			if nextRoom == farm.End {
				finished++
			}
		}

		// Print all moves for this turn
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}

package main

import "fmt"

// PrintAntMove creates the string format for an ant move
// This creates strings like "L1-room2" meaning "Ant 1 moves to room2"
func PrintAntMove(antID int, roomName string) string {
	return fmt.Sprintf("L%d-%s", antID, roomName)
}

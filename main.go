package main

import (
	"fmt"
	"os"
	"strings"
)

// main is the entry point: it reads input file, constructs the farm, finds paths, and simulates ant movements.
func main() {
	// Check if user provided exactly one argument (the filename)
	if len(os.Args) != 2 {
		fmt.Println("ERROR: usage --> go run . <filename>")
		fmt.Println("For visualization: ./lem-in <filename> | ./visualizer")
		return
	}

	filename := os.Args[1]

	// Read the input file
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR: could not read file:", err)
		return
	}

	// Parse the file content into lines
	lines := parseInput(string(content))

	// Build the farm structure from the parsed lines
	farm, err := BuildFarm(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find all possible paths from start to end
	allPaths := FindAllPaths(farm.Start, farm.End)
	if len(allPaths) == 0 {
		fmt.Println("ERROR: invalid data format, no path from ##start to ##end")
		return
	}

	// Find the best combination of paths that minimizes total moves
	best := FindOptimalPathCombination(farm.AntCount, allPaths)
	if len(best.Paths) == 0 {
		fmt.Println("ERROR: invalid data format, no valid path combination found")
		return
	}

	// Echo original input first (as required by the project)
	fmt.Print(string(content))
	// Add blank line only if content doesn't end with newline
	if !strings.HasSuffix(string(content), "\n") {
		fmt.Println()
	}
	fmt.Println()

	// Run the ant movement simulation
	RunSimulation(farm, best.Paths)
}

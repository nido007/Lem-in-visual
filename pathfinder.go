package main

import (
	"sort"
)

// FindAllPaths finds every possible route from start to end
func FindAllPaths(start, end *Room) [][]*Room {
	var result [][]*Room

	// Use depth-first search to explore all paths
	var dfs func(path []*Room, visited map[string]bool)
	dfs = func(path []*Room, visited map[string]bool) {
		current := path[len(path)-1]

		// If we reached the end, save this path
		if current == end {
			// Make a copy so we don't modify the original
			copyPath := make([]*Room, len(path))
			copy(copyPath, path)
			result = append(result, copyPath)
			return
		}

		// Try all connected rooms
		for _, neighbor := range current.Links {
			// Don't revisit rooms we've already been to
			if visited[neighbor.Name] {
				continue
			}

			// Mark this room as visited and continue exploring
			visited[neighbor.Name] = true
			dfs(append(path, neighbor), visited)
			visited[neighbor.Name] = false // Backtrack
		}
	}

	// Start the search from the starting room
	visited := map[string]bool{start.Name: true}
	dfs([]*Room{start}, visited)

	return result
}

// isCompatible checks if a path can be used alongside other paths
// Two paths are compatible if they don't share any middle rooms
func isCompatible(candidate []*Room, currentSet [][]*Room) bool {
	// Track which middle rooms are already used
	used := make(map[string]bool)
	for _, path := range currentSet {
		// Only check middle rooms (skip first and last)
		for i := 1; i < len(path)-1; i++ {
			used[path[i].Name] = true
		}
	}

	// Check if candidate path uses any already-used middle rooms
	for i := 1; i < len(candidate)-1; i++ {
		if used[candidate[i].Name] {
			return false // Conflict found
		}
	}

	return true // No conflicts
}

// FindNonOverlappingPathSets finds all possible combinations of paths that don't interfere
func FindNonOverlappingPathSets(paths [][]*Room) [][][]*Room {
	var result [][][]*Room

	// Use backtracking to find all valid combinations
	var backtrack func(start int, currentSet [][]*Room)
	backtrack = func(start int, currentSet [][]*Room) {
		// Save current combination (even if empty)
		copySet := make([][]*Room, len(currentSet))
		copy(copySet, currentSet)
		result = append(result, copySet)

		// Try adding more paths
		for i := start; i < len(paths); i++ {
			if isCompatible(paths[i], currentSet) {
				backtrack(i+1, append(currentSet, paths[i]))
			}
		}
	}

	backtrack(0, [][]*Room{})
	return result
}

// PathCombination represents a set of paths and how many turns they'll take
type PathCombination struct {
	Paths [][]*Room
	Turns int
}

// EstimateTurns calculates how many moves it will take to get all ants through
func EstimateTurns(antCount int, paths [][]*Room) int {
	if len(paths) == 0 {
		return 999999 // Infinity - no paths available
	}

	// Sort paths by length (shortest first)
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	// Distribute ants among paths
	antsLeft := antCount
	antsPerPath := make([]int, len(paths))

	// Give one ant to each path, then repeat until all ants are assigned
	for antsLeft > 0 {
		for i := range paths {
			if antsLeft == 0 {
				break
			}
			antsPerPath[i]++
			antsLeft--
		}
	}

	// Calculate maximum turns among all paths
	maxTurns := 0
	for i, path := range paths {
		if antsPerPath[i] == 0 {
			continue
		}
		// Time = path length + time for all ants to go through
		turns := (len(path) - 1) + (antsPerPath[i] - 1)
		if turns > maxTurns {
			maxTurns = turns
		}
	}

	return maxTurns
}

// FindOptimalPathCombination finds the best combination of paths
func FindOptimalPathCombination(antCount int, paths [][]*Room) PathCombination {
	combinations := FindNonOverlappingPathSets(paths)

	var best PathCombination
	best.Turns = 999999 // Start with worst case

	// Test each combination and keep the best one
	for _, combo := range combinations {
		if len(combo) == 0 {
			continue // Skip empty combinations
		}

		turns := EstimateTurns(antCount, combo)
		if turns < best.Turns {
			best = PathCombination{
				Paths: combo,
				Turns: turns,
			}
		}
	}

	return best
}

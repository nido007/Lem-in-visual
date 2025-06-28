package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Room represents a single room in the ant farm
type Room struct {
	Name     string  // The name of the room
	X, Y     int     // Position coordinates
	Links    []*Room // List of connected rooms
	Occupied bool    // Whether an ant is currently in this room
}

// Farm represents the entire ant colony
type Farm struct {
	Rooms    map[string]*Room // All rooms in the farm
	Start    *Room            // Starting room
	End      *Room            // Destination room
	AntCount int              // Number of ants to move
}

// BuildFarm reads the input and creates the farm structure
func BuildFarm(lines []string) (*Farm, error) {
	if len(lines) == 0 {
		return nil, errors.New("ERROR: invalid data format, empty input")
	}

	// First line should be the number of ants
	antCount, err := strconv.Atoi(lines[0])
	if err != nil || antCount <= 0 {
		return nil, errors.New("ERROR: invalid data format, invalid number of ants")
	}

	// Create a new farm
	farm := &Farm{
		AntCount: antCount,
		Rooms:    make(map[string]*Room),
	}

	// Flags to track special commands
	var expectStart, expectEnd bool

	// Process each line after the ant count
	for i := 1; i < len(lines); i++ {
		line := lines[i]

		// Skip empty lines
		if line == "" {
			continue
		}

		// Handle special commands and comments
		if strings.HasPrefix(line, "#") {
			if line == "##start" {
				expectStart = true
			} else if line == "##end" {
				expectEnd = true
			}
			continue
		}

		// Check if this line defines a room (has spaces)
		if strings.Contains(line, " ") {
			err := processRoom(line, farm, &expectStart, &expectEnd)
			if err != nil {
				return nil, err
			}
		} else if strings.Contains(line, "-") {
			// This line defines a tunnel between rooms
			err := processLink(line, farm)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("ERROR: invalid data format, unknown line format: %s", line)
		}
	}

	// Make sure we have both start and end rooms
	if farm.Start == nil {
		return nil, errors.New("ERROR: invalid data format, missing ##start room")
	}
	if farm.End == nil {
		return nil, errors.New("ERROR: invalid data format, missing ##end room")
	}

	return farm, nil
}

// processRoom handles a room definition line
func processRoom(line string, farm *Farm, expectStart, expectEnd *bool) error {
	// Split the line into parts: name x y
	tokens := strings.Fields(line)
	if len(tokens) != 3 {
		return fmt.Errorf("ERROR: invalid data format, malformed room: %s", line)
	}

	name, xStr, yStr := tokens[0], tokens[1], tokens[2]

	// Room names cannot start with 'L' or '#'
	if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") {
		return fmt.Errorf("ERROR: invalid data format, invalid room name: %s", name)
	}

	// Check for duplicate room names
	if _, exists := farm.Rooms[name]; exists {
		return fmt.Errorf("ERROR: invalid data format, duplicate room name: %s", name)
	}

	// Parse the coordinates
	x, err1 := strconv.Atoi(xStr)
	y, err2 := strconv.Atoi(yStr)
	if err1 != nil || err2 != nil {
		return fmt.Errorf("ERROR: invalid data format, invalid coordinates in room: %s", line)
	}

	// Create the room
	room := &Room{
		Name: name,
		X:    x,
		Y:    y,
	}

	// Add room to the farm
	farm.Rooms[name] = room

	// Set as start or end room if flagged
	if *expectStart {
		farm.Start = room
		*expectStart = false
	}
	if *expectEnd {
		farm.End = room
		*expectEnd = false
	}

	return nil
}

// processLink handles a tunnel definition line
func processLink(line string, farm *Farm) error {
	// Split the line: room1-room2
	tokens := strings.Split(line, "-")
	if len(tokens) != 2 {
		return fmt.Errorf("ERROR: invalid data format, malformed link: %s", line)
	}

	fromName, toName := tokens[0], tokens[1]

	// Prevent rooms from linking to themselves
	if fromName == toName {
		return fmt.Errorf("ERROR: invalid data format, self-linked room: %s", line)
	}

	// Make sure both rooms exist
	room1, ok1 := farm.Rooms[fromName]
	room2, ok2 := farm.Rooms[toName]
	if !ok1 || !ok2 {
		return fmt.Errorf("ERROR: invalid data format, link references unknown room(s): %s", line)
	}

	// Add bidirectional link if it doesn't already exist
	if !isLinked(room1, room2) {
		room1.Links = append(room1.Links, room2)
		room2.Links = append(room2.Links, room1)
	}

	return nil
}

// isLinked checks if two rooms are already connected
func isLinked(a, b *Room) bool {
	for _, link := range a.Links {
		if link == b {
			return true
		}
	}
	return false
}

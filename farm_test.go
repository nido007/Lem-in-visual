package main

import (
	"testing"
)

// TestBuildFarm_EmptyInput tests what happens with empty input
func TestBuildFarm_EmptyInput(t *testing.T) {
	_, err := BuildFarm([]string{})
	if err == nil {
		t.Error("BuildFarm(empty) should return error but didn't")
	}
}

// TestBuildFarm_InvalidAnts tests invalid ant counts
func TestBuildFarm_InvalidAnts(t *testing.T) {
	tests := [][]string{
		{"0"},   // Zero ants
		{"-1"},  // Negative ants
		{"abc"}, // Non-number
		{""},    // Empty
	}

	for _, test := range tests {
		_, err := BuildFarm(test)
		if err == nil {
			t.Errorf("BuildFarm(%v) should return error but didn't", test)
		}
	}
}

// TestBuildFarm_SelfLink tests rooms linking to themselves
func TestBuildFarm_SelfLink(t *testing.T) {
	lines := []string{
		"1",
		"##start",
		"A 0 0",
		"##end",
		"B 1 1",
		"A-A", // Room linking to itself
	}
	_, err := BuildFarm(lines)
	if err == nil {
		t.Error("BuildFarm(self-link) should return error but didn't")
	}
}

// TestBuildFarm_DuplicateRoom tests duplicate room names
func TestBuildFarm_DuplicateRoom(t *testing.T) {
	lines := []string{
		"1",
		"##start",
		"A 0 0",
		"##end",
		"B 1 1",
		"A 0 0", // Duplicate room name
	}
	_, err := BuildFarm(lines)
	if err == nil {
		t.Error("BuildFarm(duplicate room) should return error but didn't")
	}
}

// TestBuildFarm_MissingStartEnd tests missing start or end markers
func TestBuildFarm_MissingStartEnd(t *testing.T) {
	lines := []string{
		"1",
		"A 0 0",
		"B 1 1",
		"A-B",
	}
	_, err := BuildFarm(lines)
	if err == nil {
		t.Error("BuildFarm(missing start/end) should return error but didn't")
	}
}

// TestBuildFarm_InvalidRoomName tests invalid room names
func TestBuildFarm_InvalidRoomName(t *testing.T) {
	tests := [][]string{
		{"1", "##start", "L123 0 0", "##end", "B 1 1"}, // Starts with L
		{"1", "##start", "#bad 0 0", "##end", "B 1 1"}, // Starts with #
	}

	for _, test := range tests {
		_, err := BuildFarm(test)
		if err == nil {
			t.Errorf("BuildFarm(%v) should return error but didn't", test)
		}
	}
}

// TestBuildFarm_ValidCase tests a valid farm
func TestBuildFarm_ValidCase(t *testing.T) {
	lines := []string{
		"2",
		"##start",
		"start 0 0",
		"middle 1 1",
		"##end",
		"end 2 2",
		"start-middle",
		"middle-end",
	}

	farm, err := BuildFarm(lines)
	if err != nil {
		t.Errorf("BuildFarm(valid) returned error: %v", err)
		return
	}

	if farm.AntCount != 2 {
		t.Errorf("Expected 2 ants, got %d", farm.AntCount)
	}

	if farm.Start.Name != "start" {
		t.Errorf("Expected start room 'start', got '%s'", farm.Start.Name)
	}

	if farm.End.Name != "end" {
		t.Errorf("Expected end room 'end', got '%s'", farm.End.Name)
	}

	if len(farm.Rooms) != 3 {
		t.Errorf("Expected 3 rooms, got %d", len(farm.Rooms))
	}
}

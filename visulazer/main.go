package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Room struct {
	Name string
	X, Y int
}

type Ant struct {
	ID       int
	RoomName string
}

type Farm struct {
	Rooms    map[string]*Room
	Links    [][]string
	AntCount int
	Start    string
	End      string
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func parseInput() (*Farm, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	var movesStarted bool
	var moves []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			movesStarted = true
			continue
		}

		if movesStarted {
			moves = append(moves, line)
		} else {
			lines = append(lines, line)
		}
	}

	farm := &Farm{
		Rooms: make(map[string]*Room),
		Links: [][]string{},
	}

	var expectStart, expectEnd bool

	for i, line := range lines {
		if i == 0 {
			antCount, _ := strconv.Atoi(line)
			farm.AntCount = antCount
			continue
		}

		if strings.HasPrefix(line, "#") {
			if line == "##start" {
				expectStart = true
			} else if line == "##end" {
				expectEnd = true
			}
			continue
		}

		if strings.Contains(line, " ") {
			parts := strings.Fields(line)
			if len(parts) == 3 {
				name := parts[0]
				x, _ := strconv.Atoi(parts[1])
				y, _ := strconv.Atoi(parts[2])

				room := &Room{Name: name, X: x, Y: y}
				farm.Rooms[name] = room

				if expectStart {
					farm.Start = name
					expectStart = false
				}
				if expectEnd {
					farm.End = name
					expectEnd = false
				}
			}
		} else if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			if len(parts) == 2 {
				farm.Links = append(farm.Links, parts)
			}
		}
	}

	return farm, moves
}

func createDynamicVisualization(farm *Farm, ants map[int]*Ant) {
	fmt.Println()
	fmt.Println("        _________________")
	fmt.Println("       /                 \\")

	// Build room representations with ant markers
	room5 := "[5]"
	room3 := "[3]"
	room1 := "[1]"
	if hasAntInRoom(ants, "5") {
		room5 = "*[5]*"
	}
	if hasAntInRoom(ants, "3") {
		room3 = "*[3]*"
	}
	if hasAntInRoom(ants, "1") {
		room1 = "*[1]*"
	}

	fmt.Printf("  ____%s----%s--%s     |\n", room5, room3, room1)
	fmt.Println(" /            |    /      |")

	room6 := "[6]"
	room0 := "[0]"
	room4 := "[4]"
	if hasAntInRoom(ants, "6") {
		room6 = "*[6]*"
	}
	if hasAntInRoom(ants, "0") {
		room0 = "*[0]*"
	}
	if hasAntInRoom(ants, "4") {
		room4 = "*[4]*"
	}

	fmt.Printf("%s---%s----%s  /       |\n", room6, room0, room4)
	fmt.Println(" \\   ________/|  /        |")

	room2 := "[2]"
	room7 := "[7]"
	if hasAntInRoom(ants, "2") {
		room2 = "*[2]*"
	}
	if hasAntInRoom(ants, "7") {
		room7 = "*[7]*"
	}

	fmt.Printf("  \\ /        %s/________/\n", room2)
	fmt.Printf("  %s_________/\n", room7)
	fmt.Println()
}

func hasAntInRoom(ants map[int]*Ant, roomName string) bool {
	for _, ant := range ants {
		if ant.RoomName == roomName {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("🎨 Lem-in ASCII Art Visualizer")

	farm, moves := parseInput()

	if len(farm.Rooms) == 0 {
		fmt.Println("❌ No farm data received!")
		return
	}

	fmt.Printf("✅ Farm loaded: %d rooms, %d ants\n", len(farm.Rooms), farm.AntCount)
	fmt.Println("Press Enter to see the farm layout...")
	fmt.Scanln()

	clearScreen()
	fmt.Println("Which corresponds to the following representation:")

	ants := make(map[int]*Ant)
	createDynamicVisualization(farm, ants)

	fmt.Println("Press Enter to start animation...")
	fmt.Scanln()

	for turnNum, move := range moves {
		clearScreen()
		fmt.Printf("🐜 TURN %d: %s\n", turnNum+1, move)
		fmt.Println(strings.Repeat("=", 50))

		moveParts := strings.Fields(move)
		for _, movePart := range moveParts {
			if strings.Contains(movePart, "-") {
				parts := strings.Split(movePart, "-")
				if len(parts) == 2 && strings.HasPrefix(parts[0], "L") {
					antIDStr := strings.TrimPrefix(parts[0], "L")
					antID, err := strconv.Atoi(antIDStr)
					if err == nil {
						roomName := parts[1]
						ants[antID] = &Ant{ID: antID, RoomName: roomName}
						if roomName == farm.End {
							delete(ants, antID)
						}
					}
				}
			}
		}

		fmt.Println("Which corresponds to the following representation:")
		createDynamicVisualization(farm, ants)

		if len(ants) > 0 {
			fmt.Printf("Active ants: ")
			for id, ant := range ants {
				fmt.Printf("A%d@%s ", id, ant.RoomName)
			}
			fmt.Println()
		}

		time.Sleep(2 * time.Second)
	}

	clearScreen()
	fmt.Println("🎉 FINAL STATE:")
	fmt.Println("Which corresponds to the following representation:")
	createDynamicVisualization(farm, ants)
	fmt.Println("✨ All ants have reached their destination!")
}

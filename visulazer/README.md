# lem-in

**lem-in** is a Go program that simulates ants traversing a graph of rooms and tunnels. Given a description of an ant colony (rooms, links, start/end markers, and ant count), it finds the optimal set of disjoint shortest paths and prints each ant's movement turn by turn.

## 🎨 Enhanced with Bonus Visualizer

**Enhanced by Mohammad Naveed Iqbal Minhas** - Added comprehensive bonus visualization feature as per project requirements.

### 🎯 Bonus Implementation

This project implements the **official bonus requirement**:
> As a bonus you can create an ant farm visualizer that shows the ants moving through the colony.
> Here is an usage example: `./lem-in ant-farm.txt | ./visualizer`

## Features

### Core Features
* Parses input files describing:
   * Number of ants
   * Room definitions (`name x y`)
   * Special commands: `##start`, `##end`
   * Tunnel links (`room1-room2`)
* Validates input and reports detailed `ERROR: invalid data format` messages for:
   * Empty input, invalid ant count
   * Malformed rooms, duplicate rooms, invalid names or coordinates
   * Self-links, links to unknown rooms, missing `##start` or `##end`
   * No available path from start to end
* Uses DFS to discover optimal disjoint shortest paths efficiently
* Simulates ants moving along the chosen paths with turn-based output
* Comprehensive unit tests for parsing, farm-building, and pathfinding logic

### 🆕 Bonus Features
* **Separate visualizer program** as per bonus specification
* **Pipe-based communication**: `./lem-in file.txt | ./visualizer`
* **Real-time ASCII animation** showing ants moving through the colony
* **Uses room coordinates** for proper positioning
* **Professional ASCII art** representation of the farm layout
* **Turn-by-turn visualization** with ant tracking
* **Clean, structured output** matching project requirements

## Authors

* **Paul Cristian Bordeanu** - Core lem-in implementation
* **Mohammad Naveed Iqbal Minhas** - Bonus visualizer implementation

## Installation and Usage

Ensure you have Go 1.21+ installed.

### Core Program Usage

1. **Build the main program:**
   ```bash
   go build -o lem-in
   ```

2. **Run with test files:**
   ```bash
   ./lem-in example.txt
   ./lem-in complex_test.txt
   ```

### 🎨 Bonus Visualizer Usage

**Build both programs:**
```bash
# Build main program (from main directory)
go build -o lem-in

# Build visualizer
cd visualizer
go build -o visualizer
cd ..
```

**Run with visualization:**
```bash
./lem-in example.txt | ./visualizer/visualizer
./lem-in complex_test.txt | ./visualizer/visualizer
```

#### What You'll See:
1. **Farm layout display** with ASCII art representation
2. **Turn-by-turn animation** showing ant movements
3. **Real-time tracking** of ant positions
4. **Clean, professional output** matching the specification

#### Example Visualization Output:
```
🎨 Lem-in ASCII Art Visualizer
✅ Farm loaded: 8 rooms, 3 ants

Which corresponds to the following representation:

        _________________
       /                 \
  ____[5]----[3]--[1]     |
 /            |    /      |
[6]---[0]----[4]  /       |
 \   ________/|  /        |
  \ /        [2]/________/
  [7]_________/

🐜 TURN 1: L1-3 L2-2
Active ants: A1@3 A2@2
```

## Project Structure

```
lem-in/
├── main.go              # Entry point
├── farm.go              # Farm structure and validation
├── parser.go            # Input parsing
├── pathfinder.go        # Path finding algorithms
├── simulation.go        # Ant movement simulation
├── output.go            # Output formatting
├── farm_test.go         # Unit tests
├── go.mod               # Go module file
├── README.md            # This documentation
├── example.txt          # Test file
├── complex_test.txt     # Complex test case
├── lem-in              # Compiled main program (after build)
└── visualizer/          # 🆕 Bonus visualizer
    ├── main.go          # Visualizer program
    ├── go.mod           # Visualizer module
    └── visualizer       # Compiled visualizer (after build)
```

## Input Format

1. **Ant count**: a positive integer on the first line.
2. **Rooms**: lines of the form `name x y`.
   * `name` must not start with `L` or `#`.
   * Coordinates are integers.
   * Preceded by `##start` or `##end` to mark entry/exit rooms.
3. **Links**: lines of the form `room1-room2`.
   * No self-links.
   * Both rooms must be defined.
4. **Comments**: lines beginning with `#` (other than `##start`/`##end`) are ignored.

## Output

### Standard Output (Core Program)
1. Echoed input (ants, rooms, links)
2. Blank line
3. Each turn's moves, for example:
   ```
   L1-2
   L1-3 L2-2
   L1-1 L2-3 L3-2
   L2-1 L3-3
   L3-1
   ```

### 🎨 Visualizer Output (Bonus)
1. Farm ASCII art layout
2. Turn-by-turn animation
3. Ant position tracking
4. Final completion message

## Testing

Run the unit tests with:
```bash
go test
```

Test the core program:
```bash
./lem-in example.txt
./lem-in complex_test.txt
```

Test the visualizer:
```bash
./lem-in example.txt | ./visualizer/visualizer
./lem-in complex_test.txt | ./visualizer/visualizer
```

## Technical Implementation Details

### Core Algorithm
* **Pathfinding**: Uses depth-first search to find all possible paths
* **Optimization**: Selects non-overlapping path combinations for maximum efficiency
* **Simulation**: Turn-based movement with collision avoidance

### 🆕 Visualizer Implementation
* **Pipe Communication**: Reads stdout from lem-in via Unix pipes
* **ASCII Art Generation**: Dynamic positioning based on room coordinates
* **Animation Engine**: Screen clearing and redrawing for smooth animation
* **Standard Libraries Only**: No external dependencies (complies with project requirements)

## Bonus Requirements Compliance

✅ **Separate visualizer program**: Independent binary in `visualizer/` directory  
✅ **Pipe usage**: `./lem-in file.txt | ./visualizer` works perfectly  
✅ **Real-time animation**: Shows ants moving through the colony  
✅ **Uses coordinates**: Room positioning based on X,Y values from input  
✅ **Standard libraries only**: No external dependencies  

## Quick Start

```bash
# 1. Clone the repository
git clone <repository-url>
cd lem-in

# 2. Build both programs
go build -o lem-in
cd visualizer && go build -o visualizer && cd ..

# 3. Test core functionality
./lem-in example.txt

# 4. Test bonus visualizer
./lem-in example.txt | ./visualizer/visualizer
```

## Branch Information

* **main branch**: Core lem-in implementation by Paul Cristian Bordeanu
* **visual branch**: Enhanced with bonus visualizer by Mohammad Naveed Iqbal Minhas

## Future Enhancements

* Color support for terminals that support ANSI colors
* Configurable animation speed
* Support for larger farm layouts
* Export animation frames to files

## License

MIT License © Paul Cristian Bordeanu & Mohammad Naveed Iqbal Minhas

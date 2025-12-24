# Game of Life (Go)

A simple, terminal-based implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) written in Go.

## Features

- **Toroidal Grid**: The grid wraps around at the edges (top connects to bottom, left connects to right).
- **Customizable Size**: Set the width and height of the simulation grid via command-line flags.
- **Adjustable Speed**: Control the simulation speed by setting the delay between generations.
- **Random Start**: Initializes with a random distribution of live cells.

## Prerequisites

- [Go](https://go.dev/dl/) (version 1.22.2 or later recommended)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/ausbeldev/Game-of-Life-Go.git
   cd Game-of-Life-Go
   ```

2. Build the project:
   ```bash
   go build -o gol main.go
   ```

## Usage

Run the compiled binary:

```bash
./gol
```

To exit the simulation, press `Ctrl+C`.

### Configuration Flags

You can customize the simulation using the following flags:

| Flag | Description | Default |
|------|-------------|---------|
| `-w` | Width of the grid (columns) | 30 |
| `-h` | Height of the grid (rows) | 30 |
| `-d` | Delay between generations in milliseconds | 250 |

### Examples

Run with a larger grid:
```bash
./gol -w 80 -h 40
```

Run a faster simulation:
```bash
./gol -d 100
```

Combine flags:
```bash
./gol -w 100 -h 50 -d 50
```

## How it Works

The Game of Life is a cellular automaton devised by the British mathematician John Horton Conway in 1970.

**Rules:**
1. **Underpopulation**: Any live cell with fewer than two live neighbours dies.
2. **Survival**: Any live cell with two or three live neighbours lives on to the next generation.
3. **Overpopulation**: Any live cell with more than three live neighbours dies.
4. **Reproduction**: Any dead cell with exactly three live neighbours becomes a live cell.

This implementation renders the grid in the terminal using ANSI escape codes for smooth refreshing.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"time"
)

var directions = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func randomState(state [][]bool) {
	threshold := 0.85
	for i := range state {
		for j := range state[i] {
			state[i][j] = rand.Float64() > threshold
		}
	}
}

func drawBoard(state [][]bool, out *bufio.Writer) {
	fmt.Fprint(out, "\033[H")
	for _, row := range state {
		for _, col := range row {
			if !col {
				fmt.Fprint(out, " ")
			} else {
				fmt.Fprint(out, "█")
			}
		}
		fmt.Fprintln(out)
	}
	out.Flush()
}

func newStateCalc(state [][]bool, newState [][]bool) {
	rows := len(state)
	cols := len(state[0])
	for i, row := range state {
		for j := range row {
			count := 0
			for _, d := range directions {
				ni := (i + d[0] + rows) % rows
				nj := (j + d[1] + cols) % cols

				if state[ni][nj] {
					count++
				}
			}
			if state[i][j] {
				newState[i][j] = (count == 2 || count == 3)
			} else {
				newState[i][j] = (count == 3)
			}
		}
	}
}

func main() {
	w := flag.Int("w", 30, "width of the grid")
	h := flag.Int("h", 30, "height of the grid")
	delay := flag.Int("d", 250, "delay in ms")
	flag.Parse()

	if *w <= 0 || *h <= 0 {
		fmt.Println("Width and Height must be positive integers")
		return
	}

	state := make([][]bool, *h)
	newState := make([][]bool, *h)
	out := bufio.NewWriter(os.Stdout)
	fmt.Print("\033[2J")
	for i := range newState {
		newState[i] = make([]bool, *w)
		state[i] = make([]bool, *w)
	}
	randomState(state)
	for {
		drawBoard(state, out)
		newStateCalc(state, newState)
		state, newState = newState, state
		time.Sleep(time.Duration(*delay) * time.Millisecond)
	}
}

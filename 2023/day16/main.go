package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type dir int

const (
	up dir = iota
	down
	left
	right
)

func main() {
	f, err := os.Open("2023/day16/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file:%v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var grid [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		var row []byte
		for i := range line {
			row = append(row, line[i])
		}
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner err: %v", err)
	}

	fmt.Println("part 1 answer:", getEnergizedTiles(grid, []int{0, 0}, right))

	m := len(grid)
	n := len(grid[0])
	var maxTiles int
	// left most going right
	for row := 0; row < m; row++ {
		maxTiles = max(getEnergizedTiles(grid, []int{row, 0}, right), maxTiles)
	}
	// right most going left
	for row := 0; row < m; row++ {
		maxTiles = max(getEnergizedTiles(grid, []int{row, n - 1}, left), maxTiles)
	}
	// top going down
	for col := 0; col < n; col++ {
		maxTiles = max(getEnergizedTiles(grid, []int{0, col}, down), maxTiles)
	}
	// bottom going up
	for col := 0; col < n; col++ {
		maxTiles = max(getEnergizedTiles(grid, []int{m - 1, col}, up), maxTiles)
	}

	fmt.Println("part 2 answer:", maxTiles)
}

func getEnergizedTiles(grid [][]byte, start []int, direction dir) int {
	m := len(grid)
	n := len(grid[0])

	stack := []tileTrace{
		{
			start[0], start[1], direction,
		},
	}
	seen := map[tileTrace]bool{
		stack[0]: true,
	}
	marked := make([][]bool, m)
	for i := range marked {
		marked[i] = make([]bool, n)
	}

	var tiles int
	for len(stack) > 0 {
		tile := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !marked[tile.i][tile.j] {
			tiles++
			marked[tile.i][tile.j] = true
		}

		tr := tileTrace{
			i: tile.i,
			j: tile.j,
			d: tile.d,
		}
		switch grid[tile.i][tile.j] {
		case '.':
			switch tile.d {
			case up:
				tr.i--
			case down:
				tr.i++
			case left:
				tr.j--
			case right:
				tr.j++
			}
		case '/':
			switch tile.d {
			case up:
				tr.j++
				tr.d = right
			case down:
				tr.j--
				tr.d = left
			case right:
				tr.i--
				tr.d = up
			case left:
				tr.i++
				tr.d = down
			}
		case '\\':
			switch tile.d {
			case up:
				tr.j--
				tr.d = left
			case down:
				tr.j++
				tr.d = right
			case right:
				tr.i++
				tr.d = down
			case left:
				tr.i--
				tr.d = up
			}
		case '-':
			switch tile.d {
			case left:
				tr.j--
			case right:
				tr.j++
			case up, down:
				// take care of split on left
				split := tileTrace{tr.i, tr.j - 1, left}
				if validRow(split.i, split.j, m, n) && !seen[split] {
					stack = append(stack, split)
					seen[split] = true
				}
				tr.j++
				tr.d = right
			}
		case '|':
			switch tile.d {
			case up:
				tr.i--
			case down:
				tr.i++
			case left, right:
				// take care of split to top
				split := tileTrace{tr.i - 1, tr.j, up}
				if validRow(split.i, split.j, m, n) && !seen[split] {
					stack = append(stack, split)
					seen[split] = true
				}
				tr.i++
				tr.d = down
			}
		}
		if validRow(tr.i, tr.j, m, n) && !seen[tr] {
			stack = append(stack, tr)
			seen[tr] = true
		}
	}

	return tiles
}

func validRow(i, j, m, n int) bool {
	return i >= 0 && j >= 0 && i < m && j < n
}

type tileTrace struct {
	i int
	j int
	d dir
}

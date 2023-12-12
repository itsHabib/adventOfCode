package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("2023/day10/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file:%v", err)
	}

	scanner := bufio.NewScanner(f)
	var grid [][]byte
	start := make([]int, 2)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		var row []byte
		for i := range line {
			if line[i] == 'S' {
				start[0] = len(grid)
				start[1] = i
			}
			row = append(row, line[i])
		}
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner err: %v", err)
	}

	stack := [][][]int{{{start[0], start[1], 0}}}
	var longestPath [][]int
	var longestCycle int
	for len(stack) > 0 {
		path := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		i := path[len(path)-1][0]
		j := path[len(path)-1][1]
		length := path[len(path)-1][2]

		dirs := getDirections(grid[i][j], path[len(path)-1])
		for k := range dirs {
			newRow := dirs[k][0]
			newCol := dirs[k][1]
			if !validCell(grid, newRow, newCol) || grid[newRow][newCol] == '.' {
				continue
			}
			if grid[newRow][newCol] == 'S' {
				if length+1 > longestCycle {
					longestCycle = length + 1
					longestPath = make([][]int, len(path))
					copy(longestPath, path)
				}
				continue
			}
			if inPath(path, newRow, newCol) {
				continue
			}
			newPath := make([][]int, len(path))
			copy(newPath, path)
			newPath = append(newPath, []int{newRow, newCol, length + 1})
			stack = append(stack, newPath)
		}
	}

	fmt.Println("part 1 answer:", longestCycle/2)
	fmt.Println("part 2 answer:", countEnclosedArea(grid, longestPath))
}

func printGrid(grid [][]byte) {
	for i := range grid {
		for j := range grid[i] {
			fmt.Printf("%s ", string(grid[i][j]))
		}
		fmt.Println()
	}
}

func inPath(path [][]int, row, column int) bool {
	for i := range path {
		if path[i][0] == row && path[i][1] == column {
			return true
		}
	}
	return false
}

func validCell(grid [][]byte, row, column int) bool {
	return row >= 0 && column >= 0 && row < len(grid) && column < len(grid[0])
}

func getDirections(ch byte, coord []int) [][]int {
	switch ch {
	case '|':
		return [][]int{
			{
				coord[0] + 1, coord[1],
			},
			{
				coord[0] - 1, coord[1],
			},
		}
	case '-':
		return [][]int{
			{
				coord[0], coord[1] + 1,
			},
			{
				coord[0], coord[1] - 1,
			},
		}
	case 'L':
		return [][]int{
			{
				coord[0] - 1, coord[1],
			},
			{
				coord[0], coord[1] + 1,
			},
		}
	case 'J':
		return [][]int{
			{
				coord[0] - 1, coord[1],
			},
			{
				coord[0], coord[1] - 1,
			},
		}
	case '7':
		return [][]int{
			{
				coord[0] + 1, coord[1],
			},
			{
				coord[0], coord[1] - 1,
			},
		}

	case 'F':
		return [][]int{
			{
				coord[0] + 1, coord[1],
			},
			{
				coord[0], coord[1] + 1,
			},
		}
	case 'S':
		return [][]int{
			{
				coord[0], coord[1] + 1,
			},
			{
				coord[0], coord[1] - 1,
			},
			{
				coord[0] + 1, coord[1],
			},
			{
				coord[0] - 1, coord[1],
			},
		}
	}

	return [][]int{}
}

func countEnclosedArea(grid [][]byte, longestPath [][]int) int {
	enclosedArea := 0
	rows, cols := len(grid), len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !inPath(longestPath, i, j) && isInsideLoop(grid, i, j, longestPath) {
				enclosedArea++
			}
		}
	}

	return enclosedArea
}

func isInsideLoop(grid [][]byte, i, j int, longestPath [][]int) bool {
	count := 0
	curJ := j

	for curJ < len(grid[0]) {
		if inPath(longestPath, i, curJ) && (grid[i][curJ] == 'L' || grid[i][curJ] == 'J' || grid[i][curJ] == '|') {
			count++
		}
		curJ++
	}
	return count%2 != 0
}

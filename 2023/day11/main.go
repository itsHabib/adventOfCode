package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("2023/day11/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var grid [][]int
	var galaxy int = 1
	var emptyRows []int

	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		var seenGalaxy bool
		for i := range line {
			curGalaxy := -1
			if line[i] == '#' {
				curGalaxy = galaxy
				seenGalaxy = true
				galaxy += 1
			}
			row = append(row, curGalaxy)
		}
		if !seenGalaxy {
			emptyRows = append(emptyRows, len(grid))
		}
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error during scan: %v", err)
	}

	var emptyColumns []int
	for j := 0; j < len(grid[0]); j++ {
		empty := true
		for i := 0; i < len(grid); i++ {
			if grid[i][j] != -1 {
				empty = false
				break
			}
		}
		if empty {
			emptyColumns = append(emptyColumns, j)
		}
	}

	galaxies := make(map[int][]int)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != -1 {
				galaxies[grid[i][j]] = []int{i, j}
			}
		}
	}

	distances := make(map[int]map[int]int)

	dirs := [][]int{
		{
			0, 1,
		},
		{
			0, -1,
		},
		{
			1, 0,
		},
		{
			-1, 0,
		},
	}
	for i := 1; i < galaxy; i++ {
		queue := [][]int{galaxies[i]}
		seen := make([][]bool, len(grid))
		for idx := range seen {
			seen[idx] = make([]bool, len(grid[idx]))
		}
		seen[galaxies[i][0]][galaxies[i][1]] = true
		var distance int
		for len(queue) > 0 {
			size := len(queue)
			for j := 0; j < size; j++ {
				coord := queue[0]
				queue = queue[1:]
				cell := grid[coord[0]][coord[1]]
				if cell != -1 && distance > 0 {
					if _, ok := distances[cell]; !ok {
						m, ok := distances[i]
						if !ok {
							m = make(map[int]int)
						}
						m[cell] = distance
						distances[i] = m
					}
				}
				for k := range dirs {
					newRow := dirs[k][0] + coord[0]
					newCol := dirs[k][1] + coord[1]
					if newRow < 0 || newCol < 0 || newRow >= len(grid) || newCol >= len(grid[0]) || seen[newRow][newCol] {
						continue
					}
					queue = append(queue, []int{newRow, newCol})
					seen[newRow][newCol] = true
				}
			}
			distance++
		}
	}

	part1 := getDistanceSum(
		2,
		distances,
		emptyRows,
		emptyColumns,
		galaxies,
	)
	part2 := getDistanceSum(
		1000000,
		distances,
		emptyRows,
		emptyColumns,
		galaxies,
	)

	fmt.Println("part 1 answer:", part1)
	fmt.Println("part 2 answer:", part2)
}

func getDistanceSum(expansionFactor int, distances map[int]map[int]int, emptyRows, emptyColumns []int, galaxies map[int][]int) int {
	var sum int
	for galaxy, otherMap := range distances {
		for otherGalaxy, distance := range otherMap {
			var additional int
			for i := range emptyRows {
				if isBetween(galaxies[galaxy][0], galaxies[otherGalaxy][0], emptyRows[i]) {
					additional += expansionFactor - 1
				}
			}

			for i := range emptyColumns {
				if isBetween(galaxies[galaxy][1], galaxies[otherGalaxy][1], emptyColumns[i]) {
					additional += expansionFactor - 1
				}
			}

			sum += distance + additional
		}
	}

	return sum
}

func printGrid(grid [][]int) {
	for i := range grid {
		for j := range grid[i] {
			switch grid[i][j] {
			case -1:
				fmt.Print(". ")
			default:
				fmt.Print(grid[i][j], " ")
			}
		}
		fmt.Println()
	}
}

func isBetween(a, b, empty int) bool {
	return (empty > a && empty < b) || (empty > b && empty < a)
}

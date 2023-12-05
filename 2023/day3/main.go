package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	part1()
}

func isSymbol(s string) bool {
	// dont need to worry about letters
	return s != "." && (s < "0" || s > "9")
}

func isValidCoord[T any](grid [][]T, row, col int) bool {
	return row != -1 && col != -1 && row < len(grid) && col < len(grid[0])
}

func part1() {
	f, err := os.Open("2023/day3/input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	scanner := bufio.NewScanner(f)
	var grid [][]string
	for scanner.Scan() {
		var row []string
		line := scanner.Text()
		for i := range line {
			row = append(row, string(line[i]))
		}
		grid = append(grid, row)
	}

	for i := range grid {
		fmt.Println(grid[i])
	}

	coords := [][]int{
		{
			0, 1,
		},
		{
			1, 0,
		},
		{
			0, -1,
		},
		{
			-1, 0,
		},
		{
			1, 1,
		},
		{
			-1, 1,
		},
		{
			-1, -1,
		},
		{
			1, -1,
		},
	}

	type c struct {
		i, j int
	}
	gears := make(map[c][]int)

	var sum int
	for i := range grid {
		var j int
		for j < len(grid[i]) {
			var (
				num          string
				isPartNumber bool
				seenGear     bool
				gear         c
			)
			for j < len(grid[i]) && isDigit(grid[i][j]) {
				num += grid[i][j]
				if isPartNumber {
					j++
					continue
				}
				for k := range coords {
					coord := coords[k]
					newRow := i + coord[1]
					newCol := j + coord[0]
					if !isValidCoord(grid, newRow, newCol) {
						continue
					}
					if grid[newRow][newCol] == "*" {
						isPartNumber = true
						seenGear = true
						gear = c{newRow, newCol}
						break
					}
					if isSymbol(grid[newRow][newCol]) {
						isPartNumber = true
						break
					}
				}
				j++
			}
			if num != "" && isPartNumber {
				n, _ := strconv.Atoi(num)
				sum += n
				if seenGear {
					gears[gear] = append(gears[gear], n)
				}
			}
			j++
		}
	}

	fmt.Println("part 1 answer:", sum)

	var ratioSum int
	for _, v := range gears {
		if len(v) != 2 {
			continue
		}
		ratioSum += v[0] * v[1]
	}

	fmt.Println("part 2 answer:", ratioSum)
}

func isDigit(s string) bool {
	return s >= "0" && s <= "9"
}

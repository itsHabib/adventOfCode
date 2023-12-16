package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("2023/day14/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file:%v", err)
	}
	defer f.Close()

	var grid [][]byte
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var row []byte
		line := scanner.Text()
		for i := range line {
			row = append(row, line[i])
		}
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner err:%v", err)
	}

	rollNorth(grid)
	ans := getLoad(grid)

	fmt.Println("part 1 answer:", ans)

	var count int
	scores := make(map[int][]int)
	target := 1000000000
	for count < target {
		grid = rollCycle(grid)
		load := getLoad(grid)
		count++
		counts := scores[load]
		counts = append(counts, count)
		scores[load] = counts
		if len(counts) >= 4 {
			lastCycle := counts[len(counts)-1] - counts[len(counts)-2]
			if lastCycle == counts[len(counts)-2]-counts[len(counts)-3] {
				count = target - (target-count)%lastCycle
			}
		}
	}
	fmt.Println("part 2 answer:", getLoad(grid))
}

func getLoad(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	var ans int
	for i := 0; i < m; i++ {
		var count int
		for j := 0; j < n; j++ {
			if grid[i][j] == 'O' {
				count++
			}
		}
		if count > 0 {
			ans += count * (m - i)
		}
	}

	return ans
}

func rollCycle(grid [][]byte) [][]byte {
	for i := 0; i < 4; i++ {
		rollNorth(grid)
		grid = rotateClockWise(grid)
	}

	return grid
}

func rollNorth(grid [][]byte) {
	m := len(grid)
	n := len(grid[0])

	for j := 0; j < n; j++ {
		var lastZero = -1
		for i := m - 1; i > 0; i-- {
			if grid[i][j] == 'O' {
				lastZero = max(i, lastZero)
			}
			if grid[i][j] == 'O' && grid[i-1][j] == '.' {
				grid[i][j], grid[i-1][j] = grid[i-1][j], grid[i][j]
				for k := i; k < len(grid) && k < lastZero+1; k++ {
					i++
				}
			}

		}
	}
}

func rotateClockWise(grid [][]byte) [][]byte {
	m := len(grid)
	n := len(grid[0])
	rotated := make([][]byte, n)

	for i := range rotated {
		rotated[i] = make([]byte, m)
		for j := range rotated[i] {
			rotated[i][j] = grid[m-j-1][i]
		}
	}

	return rotated
}

func printGrid(grid [][]byte) {
	for i := range grid {
		for j := range grid[i] {
			fmt.Printf("%s ", string(grid[i][j]))
		}
		fmt.Println()
	}
}

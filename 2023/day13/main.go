package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("2023/day13/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var grids [][][]byte
	var currentGrid [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			grids = append(grids, currentGrid)
			currentGrid = [][]byte{}
			continue
		}
		currentGrid = append(currentGrid, []byte(line))
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner err: %v", err)
	}
	if len(currentGrid) > 0 {
		grids = append(grids, currentGrid)
	}

	var ans int
	for i := range grids {
		ans += calculateSymmetry(grids[i], false)
	}
	fmt.Println("part 1 answer:", ans)
	var p2Ans int
	for i := range grids {
		p2Ans += calculateSymmetry(grids[i], true)
	}
	fmt.Println("part 2 answer:", p2Ans)

}

func calculateSymmetry(grid [][]byte, pt2 bool) int {
	m := len(grid)
	n := len(grid[0])
	var ans int

	for i := 0; i < m-1; i++ {
		var misses int
		for j := 0; j < m; j++ {
			up := i - j
			down := i + j + 1
			if up < 0 || down >= m {
				continue
			}
			for k := 0; k < n; k++ {
				if grid[up][k] != grid[down][k] {
					misses++
				}
			}
		}
		if pt2 && misses == 1 || !pt2 && misses == 0 {
			ans += (i + 1) * 100
		}
	}

	for i := 0; i < n-1; i++ {
		var misses int
		for j := 0; j < n; j++ {
			left := i - j
			right := i + j + 1
			if left < 0 || right >= n {
				continue
			}
			for k := 0; k < m; k++ {
				if grid[k][left] != grid[k][right] {
					misses++
				}
			}
		}
		if pt2 && misses == 1 || !pt2 && misses == 0 {
			ans += i + 1
		}
	}

	return ans
}

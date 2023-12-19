package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	f, err := os.Open("2023/day17/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file:%v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var grid [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for i := range line {
			row = append(row, int(line[i]-'0'))
		}
		grid = append(grid, row)
	}
	fmt.Println("part 1 answer:", minHeatLoss(grid, 1, 3))
	fmt.Println("part 2 answer:", minHeatLoss(grid, 4, 10))
}

func minHeatLoss(grid [][]int, minSteps, maxSteps int) int {
	var q pq
	heap.Init(&q)
	start := cell{
		point: point{
			i: 0,
			j: 0,
		},
		weight:   0,
		dir:      right,
		dirCount: 0,
	}
	heap.Push(&q, start)
	start.dir = down
	heap.Push(&q, start)

	m := len(grid)
	n := len(grid[0])
	dirs := [][]int{
		{
			0, 1, int(right),
		},
		{
			1, 0, int(down),
		},
		{
			0, -1, int(left),
		},
		{
			-1, 0, int(up),
		},
	}
	seen := make(map[[4]int]bool)
	seen[[4]int{0, 0, int(right), 0}] = true
	seen[[4]int{0, 0, int(down), 0}] = true

	var minWeight = math.MaxInt
	for q.Len() > 0 {
		c := heap.Pop(&q).(cell)
		if c.point.i == m-1 && c.point.j == n-1 && c.dirCount >= minSteps {
			return c.weight
		}

		for i := range dirs {
			np := point{dirs[i][0] + c.point.i, dirs[i][1] + c.point.j}
			if !validPoint(np, m, n) || c.dir.reverse() == direction(dirs[i][2]) || (c.dir != direction(dirs[i][2]) && c.dirCount < minSteps) {
				continue
			}
			dirCount := 1
			if c.dir == direction(dirs[i][2]) {
				dirCount = c.dirCount + 1
			}
			if dirCount > maxSteps {
				continue
			}
			if seen[[4]int{np.i, np.j, dirs[i][2], dirCount}] {
				continue
			}
			newDist := c.weight + grid[np.i][np.j]
			nc := cell{
				point:    np,
				weight:   newDist,
				dirCount: dirCount,
				dir:      direction(dirs[i][2]),
			}
			heap.Push(&q, nc)
			seen[[4]int{np.i, np.j, dirs[i][2], dirCount}] = true
		}
	}

	return minWeight
}

func validPoint(p point, m, n int) bool {
	return p.i >= 0 && p.j >= 0 && p.i < m && p.j < n
}

type direction int

func (d direction) reverse() direction {
	switch d {
	case up:
		return down
	case left:
		return right
	case right:
		return left
	case down:
		return up
	}

	return -1
}

func (d direction) String() string {
	switch d {
	case up:
		return "up"
	case left:
		return "left"
	case right:
		return "right"
	case down:
		return "down"
	}

	return ""
}

const (
	up direction = iota
	down
	left
	right
)

type point struct {
	i int
	j int
}

type cell struct {
	point    point
	weight   int
	dir      direction
	dirCount int
}

type pq []cell

func (p pq) Len() int           { return len(p) }
func (p pq) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p pq) Less(i, j int) bool { return p[i].weight < p[j].weight }
func (p *pq) Push(x any) {
	*p = append(*p, x.(cell))
}
func (p *pq) Pop() any {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]

	return x
}

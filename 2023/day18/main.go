package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open("2023/day18/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	start := point{0, 0}
	path := []point{start}

	for sc.Scan() {
		parts := strings.Split(sc.Text(), " ")
		dir := direction(parts[0])
		steps, _ := strconv.Atoi(parts[1])
		for i := 0; i < steps; i++ {
			start = start.step(dir, 1)
			path = append(path, start)
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scanner err: %v", err)
	}

	fmt.Println("part 1 answer:", int(shoeLace(path))+(len(path)/2)+1)
}

func part2() {
	f, err := os.Open("2023/day18/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	start := point{0, 0}
	path := []point{start}

	for sc.Scan() {
		parts := strings.Split(sc.Text(), " ")
		stepsHex := parts[2][2 : len(parts[2])-2]
		steps, _ := strconv.ParseInt(stepsHex, 16, 64)
		dir := part2Dir(parts[2][len(parts[2])-2])
		for i := 0; i < int(steps); i++ {
			start = start.step(dir, 1)
			path = append(path, start)
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scanner err: %v", err)
	}

	fmt.Println("part 2 answer:", int(shoeLace(path))+(len(path)/2)+1)
}

func part2Dir(ch byte) direction {
	switch ch {
	case '0':
		return right
	case '1':
		return down
	case '2':
		return left
	case '3':
		return up
	}

	return ""
}

func shoeLace(path []point) float64 {
	area := 0
	n := len(path)

	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += path[i].i * path[j].j
		area -= path[j].i * path[i].j
	}

	return math.Abs(float64(area)) / 2.0
}

type point struct {
	i int
	j int
}

func (p point) step(d direction, steps int) point {
	switch d {
	case up:
		return point{p.i - steps, p.j}
	case down:
		return point{p.i + steps, p.j}
	case left:
		return point{p.i, p.j - steps}
	case right:
		return point{p.i, p.j + steps}
	}

	return point{}
}

type direction string

const (
	up    direction = "U"
	down  direction = "D"
	left  direction = "L"
	right direction = "R"
)

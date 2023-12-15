package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open("2023/day8/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	directions := scanner.Text()

	type nodeDirections struct {
		left  string
		right string
	}

	nodes := make(map[string]nodeDirections)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, " = ")
		val := parts[0]
		lr := strings.Split(parts[1], ",")
		left := strings.TrimPrefix(lr[0], "(")
		right := strings.TrimPrefix(strings.TrimSuffix(lr[1], ")"), " ")
		nodes[val] = nodeDirections{left: left, right: right}
	}

	next := "AAA"
	var (
		i     int
		steps int
	)
	for {
		if next == "ZZZ" {
			break
		}
		n := nodes[next]
		switch directions[i] {
		case 'L':
			next = n.left
		case 'R':
			next = n.right
		}
		steps++
		i = (i + 1) % len(directions)
	}

	fmt.Println("part 1 answer:", steps)
}

type node struct {
	val   string
	left  string
	right string
}

func part2() {
	f, err := os.Open("2023/day8/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	directions := scanner.Text()

	nodes := make(map[string]node)
	var endAs []node
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, " = ")
		val := parts[0]
		lr := strings.Split(parts[1], ",")
		left := strings.TrimPrefix(lr[0], "(")
		right := strings.TrimPrefix(strings.TrimSuffix(lr[1], ")"), " ")
		n := node{val: val, left: left, right: right}
		nodes[val] = n
		if strings.HasSuffix(val, "A") {
			endAs = append(endAs, n)
		}
	}

	var stepCounts []int
	for i := range endAs {
		stepCounts = append(stepCounts, getStepCount(directions, endAs[i], nodes))
	}
	ans := stepCounts[0]
	for i := range stepCounts[1:] {
		ans = lcm(ans, stepCounts[1:][i])
	}

	fmt.Println("part 2 answer:", ans)
}

func getStepCount(directions string, next node, nodes map[string]node) int {
	var (
		steps int
		i     int
	)
	for {
		if strings.HasSuffix(next.val, "Z") {
			return steps
		}
		switch directions[i] {
		case 'L':
			next = nodes[next.left]
		case 'R':
			next = nodes[next.right]
		}
		steps++
		i = (i + 1) % len(directions)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

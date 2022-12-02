package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	inputFile, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %s", err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var (
		elf    int
		maxCal int
	)
	counter := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()

		// new elf
		if line == "" {
			elf++
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("cant convert line to num: %s", err)
		}
		counter[elf] += num

		maxCal = max(maxCal, counter[elf])
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %s", err)
	}

	fmt.Printf("part 1: maxCal=%d\n", maxCal)
}

func part2() {
	inputFile, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %s", err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var elf int
	// arbitrary number of elves
	counter := make([]int, 1000)
	for scanner.Scan() {
		line := scanner.Text()

		// new elf
		if line == "" {
			elf++
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("cant convert line to num: %s", err)
		}
		counter[elf] += num
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %s", err)
	}

	sort.Ints(counter)
	var maxCal int
	for i := len(counter) - 1; i >= len(counter)-3; i-- {
		maxCal += counter[i]
	}

	fmt.Printf("part 2: maxCal=%d\n", maxCal)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

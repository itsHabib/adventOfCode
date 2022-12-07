package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	part1()
	part2()
}

func part1() {
	inputFile, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %s", err)
	}
	defer inputFile.Close()

	var priority int
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		firstHalf := line[:len(line)/2]
		secondHalf := line[len(line)/2:]
		seen := make(map[rune]struct{})
		for i := range firstHalf {
			seen[rune(firstHalf[i])] = struct{}{}
		}
		var item rune
		for i := range secondHalf {
			if _, ok := seen[rune(secondHalf[i])]; ok {
				item = rune(secondHalf[i])
				break
				// assuming this will always work
			}
		}
		if unicode.IsUpper(item) {
			priority += int(item - 38)
		} else {
			priority += int(item - 96)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error during scan: %s", err)
	}

	fmt.Printf("part 1=%d\n", priority)
}

func part2() {
	inputFile, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %s", err)
	}
	defer inputFile.Close()

	var (
		priority   int
		lineNumber int
	)

	scanner := bufio.NewScanner(inputFile)
	seen := make(map[rune]int)
	for scanner.Scan() {
		line := scanner.Text()
		switch lineNumber {
		case 0:
			for i := range line {
				seen[rune(line[i])] = lineNumber + 1
			}
			lineNumber++
		case 1:
			for i := range line {
				if _, ok := seen[rune(line[i])]; ok {
					seen[rune(line[i])] = lineNumber + 1
				}
			}
			lineNumber++
		case 2:
			for i := range line {
				if count, ok := seen[rune(line[i])]; ok && count == 2 {
					seen[rune(line[i])] = lineNumber + 1
				}
			}
			// find item with count of 2 - 0 indexed
			for k, v := range seen {
				if v == 3 {
					if unicode.IsUpper(k) {
						priority += int(k - 38)
					} else {
						priority += int(k - 96)
					}
					break
				}
			}
			seen = make(map[rune]int)
			lineNumber = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error during scan: %s", err)
	}

	fmt.Printf("part 2=%d\n", priority)
}

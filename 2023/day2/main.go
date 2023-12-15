package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open("2023/day2/input.txt")
	if err != nil {
		log.Fatalf("unable to get input file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var sum int
	for scanner.Scan() {
		text := scanner.Text()
		gameParts := strings.Split(text, ":")
		idStr := strings.Split(gameParts[0], " ")[1]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatalf("not an int: %s", idStr)
		}
		valid := true
		gameSets := strings.Split(gameParts[1], ";")
		for i := range gameSets {
			var (
				remainingRed   = 12
				remainingGreen = 13
				remainingBlue  = 14
			)
			pulls := strings.Split(gameSets[i], ",")
			for j := range pulls {
				pull := strings.Split(pulls[j], " ")
				count, err := strconv.Atoi(pull[1])
				if err != nil {
					log.Fatalf("not an int: %s", pull[1])
				}
				color := pull[2]
				switch color {
				case "red":
					remainingRed -= count
				case "blue":
					remainingBlue -= count
				case "green":
					remainingGreen -= count
				}
			}
			if remainingBlue < 0 || remainingRed < 0 || remainingGreen < 0 {
				valid = false
				break
			}
		}
		if valid {
			sum += id
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner err: %v", err)
	}

	fmt.Println("part 1 answer:", sum)
}

func part2() {
	f, err := os.Open("2023/day2/input.txt")
	if err != nil {
		log.Fatalf("unable to get input file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var sum int
	for scanner.Scan() {
		text := scanner.Text()
		gameParts := strings.Split(text, ":")
		var (
			maxRed   = 0
			maxGreen = 0
			maxBlue  = 0
		)
		gameSets := strings.Split(gameParts[1], ";")
		for i := range gameSets {
			pulls := strings.Split(gameSets[i], ",")
			for j := range pulls {
				pull := strings.Split(pulls[j], " ")
				count, err := strconv.Atoi(pull[1])
				if err != nil {
					log.Fatalf("not an int: %s", pull[1])
				}
				color := pull[2]
				switch color {
				case "red":
					maxRed = max(maxRed, count)
				case "blue":
					maxBlue = max(maxBlue, count)
				case "green":
					maxGreen = max(maxGreen, count)
				}
			}
		}
		sum += maxRed * maxGreen * maxBlue
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner err: %v", err)
	}

	fmt.Println("part 2 answer:", sum)
}

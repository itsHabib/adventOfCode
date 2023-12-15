package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open("2023/day6/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var times []int
	var distances []int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Time:") {
			times = collectNums(strings.Split(strings.TrimSpace(line[5:]), " "))
			continue
		}
		if strings.HasPrefix(line, "Distance:") {
			distances = collectNums(strings.Split(strings.TrimSpace(line[9:]), " "))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error while scanning: %v", err)
	}

	var mult int = 1
	for i := 0; i < len(times); i++ {
		winCount := getWinCount(times[i], distances[i])
		if winCount > 0 {
			mult *= winCount
		}
	}

	fmt.Println("part 1 answer:", mult)
}

func part2() {
	f, err := os.Open("2023/day6/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var (
		time     int
		distance int
	)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Time:") {
			var timeStr strings.Builder
			s := line[5:]
			for i := range s {
				if unicode.IsDigit(rune(s[i])) {
					timeStr.WriteByte(s[i])
				}
			}
			time, _ = strconv.Atoi(timeStr.String())
			continue
		}
		if strings.HasPrefix(line, "Distance:") {
			var distStr strings.Builder
			s := line[5:]
			for i := range s {
				if unicode.IsDigit(rune(s[i])) {
					distStr.WriteByte(s[i])
				}
			}
			distance, _ = strconv.Atoi(distStr.String())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error while scanning: %v", err)
	}

	fmt.Println("part 2 answer:", getWinCount(time, distance))
}

func getWinCount(time, distance int) int {
	var winCount int
	for j := 1; j < time; j++ {
		dist := j * (time - j)
		if dist > distance {
			winCount++
			continue
		}
		if winCount > 0 {
			break
		}
	}

	return winCount
}

func collectNums(parts []string) []int {
	var nums []int
	for i := range parts {
		if parts[i] == "" {
			continue
		}
		n, _ := strconv.Atoi(parts[i])
		nums = append(nums, n)
	}

	return nums
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	ROCK     int = 1
	PAPER    int = 2
	SCISSORS int = 3
)

func main() {
	part1()
	part2()
}

func part1() {
	inputFile, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %s", err)
	}
	defer inputFile.Close()

	var score int
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		attempts := strings.Split(line, " ")
		if len(attempts) != 2 {
			log.Fatalf("unexpected line: %s", line)
		}
		score += getWinningScore(getPlayerAttempt(attempts[0]), getStrategyAttempt(attempts[1]))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error during scan: %s", err)
	}

	fmt.Printf("part 1 score=%d\n", score)
}

func part2() {
	inputFile, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %s", err)
	}
	defer inputFile.Close()

	var score int
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		attempts := strings.Split(line, " ")
		if len(attempts) != 2 {
			log.Fatalf("unexpected line: %s", line)
		}
		playerAttempt := getPlayerAttempt(attempts[0])
		score += getWinningScore(playerAttempt, getStrategyAttemptPt2(playerAttempt, attempts[1]))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error during scan: %s", err)
	}

	fmt.Printf("part 2 score=%d\n", score)
}

func getStrategyAttempt(str string) int {
	switch str {
	case "X":
		return ROCK
	case "Y":
		return PAPER
	case "Z":
		return SCISSORS
	default:
		log.Fatalf("unexpected str: %s", str)
	}

	return 0
}

func getPlayerAttempt(str string) int {
	switch str {
	case "A":
		return ROCK
	case "B":
		return PAPER
	case "C":
		return SCISSORS
	default:
		log.Fatalf("unexpected str: %s", str)
	}

	return 0
}

// assuming a is other player and b is us
func getWinningScore(a, b int) int {
	if a == b {
		return 3 + a
	}

	if (a == ROCK && b == SCISSORS) || (a == PAPER && b == ROCK) || (a == SCISSORS && b == PAPER) {
		return 0 + b
	}

	return 6 + b
}

func getStrategyAttemptPt2(playerAttempt int, str string) int {
	switch str {
	case "X":
		// losing
		switch playerAttempt {
		case ROCK:
			return SCISSORS
		case PAPER:
			return ROCK
		case SCISSORS:
			return PAPER
		default:
			log.Fatalf("unexpected player attempt: %d", playerAttempt)
		}
	case "Y":
		return playerAttempt
	case "Z":
		// winning
		switch playerAttempt {
		case ROCK:
			return PAPER
		case PAPER:
			return SCISSORS
		case SCISSORS:
			return ROCK
		default:
			log.Fatalf("unexpected player attempt: %d", playerAttempt)
		}
	default:
		log.Fatalf("unexpected str: %s", str)
	}

	return 0
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("2023/day4/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}

	scanner := bufio.NewScanner(f)
	var sum int
	var card int
	var didntWin int
	counts := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		var i int
		for i < len(line) && line[i] != ':' {
			i++
		}
		parts := strings.Split(line[i+1:], "|")
		winningCards := strings.Split(strings.TrimSpace(parts[0]), " ")
		playingCards := strings.Split(strings.TrimSpace(parts[1]), " ")
		var (
			winCount float64 = -1
		)
		for j := range playingCards {
			if playingCards[j] != "" && isWinningCard(winningCards, playingCards[j]) {
				winCount++
			}
		}
		if winCount != -1 {
			fmt.Println(winCount)
			score := int(math.Pow(2, winCount))
			sum += score
			counts[card]++
			for k := 0; k < counts[card]; k++ {
				for j := 1; j <= int(winCount+1); j++ {
					counts[card+j]++
				}
			}
		} else {
			didntWin++
		}
		card++
	}

	fmt.Println("part 1 answer:", sum)

	var sumCopies int
	for _, count := range counts {
		sumCopies += count
	}

	fmt.Println("part 2 answer:", sumCopies+didntWin)
}

func isWinningCard(cards []string, card string) bool {
	for i := range cards {
		if cards[i] == card {
			return true
		}
	}

	return false
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	NoScore = iota
	HighCard
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var cardRanksPt1 = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var cardRanksPt2 = map[byte]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open("2023/day7/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	type handDetail struct {
		hand  string
		score int
		wager int
	}
	var hands []handDetail
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		hand := parts[0]
		wager := parts[1]
		score := getScore(hand)
		bet, _ := strconv.Atoi(wager)
		hands = append(hands, handDetail{
			hand:  hand,
			score: score,
			wager: bet,
		})
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].score != hands[j].score {
			return hands[i].score < hands[j].score
		}
		for k := range hands[i].hand {
			if hands[i].hand[k] != hands[j].hand[k] {
				return cardRanksPt1[hands[i].hand[k]] < cardRanksPt1[hands[j].hand[k]]
			}
		}
		return false
	})

	if err := scanner.Err(); err != nil {
		log.Fatalf("err during scan: %v", err)
	}

	var winnings int
	for i := range hands {
		winnings += hands[i].wager * (i + 1)
	}

	fmt.Println("part 1 answer:", winnings)
}

func part2() {
	f, err := os.Open("2023/day7/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	type handDetail struct {
		hand  string
		score int
		wager int
	}
	var hands []handDetail
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		hand := parts[0]
		wager := parts[1]
		var jCount int
		for i := range hand {
			if hand[i] == 'J' {
				jCount++
			}
		}
		score := getScore(hand)
		if jCount > 0 {
			switch score {
			case FourOfAKind, FullHouse:
				score = FiveOfAKind
			case ThreeOfAKind:
				score = FourOfAKind
			case TwoPair:
				if jCount == 2 {
					score = FourOfAKind
					break
				}
				if jCount == 1 {
					score = FullHouse
				}
			case OnePair:
				score = ThreeOfAKind
			case HighCard:
				score = OnePair
			}
		}

		bet, _ := strconv.Atoi(wager)
		hands = append(hands, handDetail{
			hand:  hand,
			score: score,
			wager: bet,
		})
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].score != hands[j].score {
			return hands[i].score < hands[j].score
		}
		for k := range hands[i].hand {
			if hands[i].hand[k] != hands[j].hand[k] {
				return cardRanksPt2[hands[i].hand[k]] < cardRanksPt2[hands[j].hand[k]]
			}
		}
		return false
	})

	if err := scanner.Err(); err != nil {
		log.Fatalf("err during scan: %v", err)
	}

	var winnings int
	for i := range hands {
		winnings += hands[i].wager * (i + 1)
	}

	fmt.Println("part 2 answer:", winnings)
}

func getScore(hand string) int {
	counts := make(map[byte]int)
	for i := range hand {
		c := counts[hand[i]]
		c++
		counts[hand[i]] = c
	}
	score := NoScore
getScore:
	for _, count := range counts {
		switch count {
		case 5:
			score = FiveOfAKind
			break getScore
		case 4:
			score = FourOfAKind
			break getScore
		case 3:
			switch score {
			case NoScore:
				score = ThreeOfAKind
			case OnePair:
				score = FullHouse
				break getScore
			}
		case 2:
			switch score {
			case NoScore:
				score = OnePair
			case OnePair:
				score = TwoPair
				break getScore
			case ThreeOfAKind:
				score = FullHouse
				break getScore
			}
		}
	}
	if score == NoScore {
		score = HighCard
	}

	return score
}

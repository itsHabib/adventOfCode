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
	f, err := os.Open("2023/day9/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var sequences [][]int
	for scanner.Scan() {
		var nums []int
		numStrings := strings.Split(scanner.Text(), " ")
		for i := range numStrings {
			if numStrings[i] == "" {
				continue
			}
			n, _ := strconv.Atoi(numStrings[i])
			nums = append(nums, n)
		}
		sequences = append(sequences, nums)
	}

	var part1Sum int
	var part2Sum int
	for i := range sequences {
		sequence := sequences[i]
		diffs := make([]int, len(sequence))
		copy(diffs, sequence)
		nextStack := []int{sequence[len(sequence)-1]}
		prevStack := []int{sequence[0]}
		for {
			next := make([]int, len(diffs))
			copy(next, diffs)
			var zeros int
			diffs = []int{}
			for j := 0; j < len(next)-1; j++ {
				diff := next[j+1] - next[j]
				if diff == 0 {
					zeros++
				}
				diffs = append(diffs, diff)
			}
			if len(diffs) == zeros {
				break
			}
			nextStack = append(nextStack, diffs[len(diffs)-1])
			prevStack = append(prevStack, diffs[0])
		}
		var nextNum int
		for len(nextStack) > 0 {
			n := nextStack[len(nextStack)-1]
			nextStack = nextStack[:len(nextStack)-1]
			nextNum += n
		}
		part1Sum += nextNum
		nextNum = 0
		for len(prevStack) > 0 {
			n := prevStack[len(prevStack)-1]
			prevStack = prevStack[:len(prevStack)-1]
			nextNum = n - nextNum
		}
		part2Sum += nextNum
	}

	fmt.Println("part 1 answer:", part1Sum)
	fmt.Println("part 2 answer:", part2Sum)
}

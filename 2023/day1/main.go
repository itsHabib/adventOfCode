package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	part1()
	part2()
}
func part1() {
	f, err := os.Open("2023/day1/input.txt")
	if err != nil {
		log.Fatalf("unable to open input.txt: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var sum int

	for scanner.Scan() {
		line := scanner.Text()
		var nums []int
		for i := range line {
			if unicode.IsDigit(rune(line[i])) {
				nums = append(nums, int(line[i]-'0'))
				continue
			}
		}

		digit := nums[0] * 10
		second := nums[0]
		if len(nums) > 1 {
			second = nums[len(nums)-1]
		}

		sum += digit + second

	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("encountered scanner err: %v", err)
	}

	fmt.Printf("part 1 answer: %d\n", sum)
}

func part2() {
	f, err := os.Open("2023/day1/input.txt")
	if err != nil {
		log.Fatalf("unable to open input.txt: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var sum int

	for scanner.Scan() {
		line := scanner.Text()
		sum += getDigit(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("encountered scanner err: %v", err)
	}

	fmt.Printf("part 2 answer: %d\n", sum)
}

func getDigit(s string) int {
	wordStart := -1
	var nums []int
	for i := range s {
		if unicode.IsDigit(rune(s[i])) {
			nums = append(nums, int(s[i]-'0'))
			wordStart = -1
			continue
		}
		if wordStart == -1 {
			wordStart = i
		}
		if n := getWord(s[wordStart : i+1]); n != -1 {
			nums = append(nums, n)
			wordStart = i
		}
	}

	digit := nums[0] * 10
	second := nums[0]
	if len(nums) > 1 {
		second = nums[len(nums)-1]
	}

	return digit + second
}

func getWord(s string) int {
	if n, ok := wordNums[s]; ok {
		return n
	}

	for w, n := range wordNums {
		if strings.Contains(s, w) {
			return n
		}
	}

	return -1
}

var wordNums = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

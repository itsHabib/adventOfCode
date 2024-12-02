package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const input = "2024/day1/input.txt"

func main() {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}

	var (
		s      = bufio.NewScanner(f)
		rcount = make(map[int]int)
		left   []int
		right  []int
	)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, "   ")
		l, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		r, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		rcount[r]++

		left = append(left, l)
		right = append(right, r)
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(left)
	sort.Ints(right)

	var (
		part1 int
		part2 int
		i     int
	)
	for i < len(left) && i < len(right) {
		l := left[i]
		r := right[i]

		part1 += int(math.Abs(float64(r - l)))
		part2 += l * rcount[l]
		i++
	}

	fmt.Println("part 1 answer:", part1)
	fmt.Println("part 2 answer:", part2)
}

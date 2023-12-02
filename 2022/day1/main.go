package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type ElfCals struct {
	elf  int
	cals int
}

func main() {

	elf, cal, err := mostCalories()
	if err != nil {
		log.Fatalf("could not get most calories: %v", err)
	}
	fmt.Println("most cals", cal, "elf", elf)
}

func mostCalories() (int, int, error) {
	f, err := os.Open("day1/input.txt")
	if err != nil {
		return 0, 0, fmt.Errorf("could not open input file: %w", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	m := make(map[int]int)
	var cur int
	for sc.Scan() {
		txt := sc.Text()
		if txt == "" {
			cur++
			continue
		}
		num, err := strconv.Atoi(txt)
		if err != nil {
			return 0, 0, fmt.Errorf("could not convert string to int: %w", err)
		}
		m[cur] += num
	}
	if sc.Err() != nil {
		return 0, 0, fmt.Errorf("could not scan input file: %w", err)
	}

	var (
		maxCal int
		elf    = -1
	)
	for e, v := range m {
		if v > maxCal {
			maxCal = v
			elf = e
		}
	}

	return elf, maxCal, nil
}

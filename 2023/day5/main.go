package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type mapper struct {
	source int
	dest   int
	length int
}

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open("2023/day5/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}

	scanner := bufio.NewScanner(f)

	var seeds []int
	var gardenMapper [][]mapper
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "seeds: ") {
			nums := strings.Split(line[7:], " ")
			for i := range nums {
				n, _ := strconv.Atoi(nums[i])
				seeds = append(seeds, n)
			}
			continue
		}
		if strings.Contains(line, "map:") {
			var maps []mapper
			for {
				scanner.Scan()
				l := scanner.Text()
				if l == "" {
					break
				}
				maps = append(maps, getMapRange(scanner.Text()))
			}
			gardenMapper = append(gardenMapper, maps)
		}
	}

	fmt.Println("part 1 answer:", getMinLocation(gardenMapper, seeds))
}

func part2() {
	n := time.Now()
	f, err := os.Open("2023/day5/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}

	scanner := bufio.NewScanner(f)

	var seeds []mapper
	var gardenMapper [][]mapper
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "seeds: ") {
			nums := strings.Split(line[7:], " ")
			for i := 0; i < len(nums)-1; i += 2 {
				start, _ := strconv.Atoi(nums[i])
				length, _ := strconv.Atoi(nums[i+1])
				seeds = append(seeds, mapper{
					source: start,
					dest:   start + length - 1,
					length: length,
				})
			}
			continue
		}
		if strings.Contains(line, "map:") {
			var maps []mapper
			for {
				scanner.Scan()
				l := scanner.Text()
				if l == "" {
					break
				}
				maps = append(maps, getMapRange(scanner.Text()))
			}
			gardenMapper = append(gardenMapper, maps)
		}
	}

	minLocation := math.MaxInt
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	seen := make(map[mapper]int)
	for i := range seeds {
		seedRange := seeds[i]
		wg.Add(1)
		go func(sr mapper) {
			for j := sr.source; j < sr.dest; j++ {
				location := getLocation(gardenMapper, j)
				mu.Lock()
				if location < minLocation {
					minLocation = location
					seen[sr] = minLocation
				}
				mu.Unlock()
			}
			wg.Done()
		}(seedRange)
	}

	wg.Wait()
	fmt.Println("part 2 answer:", minLocation, "taken:", time.Since(n).Seconds(), "s")
}

func getMapRange(s string) mapper {
	items := strings.Split(s, " ")
	var nums []int
	for j := range items {
		n, _ := strconv.Atoi(items[j])
		nums = append(nums, n)
	}
	return mapper{
		dest:   nums[0],
		source: nums[1],
		length: nums[2],
	}
}

func getMinLocation(gm [][]mapper, seeds []int) int {
	minLocation := math.MaxInt
	for i := range seeds {
		minLocation = min(minLocation, getLocation(gm, seeds[i]))
	}

	return minLocation
}

func getLocation(gm [][]mapper, seed int) int {
	chk := seed
	for i := range gm {
		maps := gm[i]
		for j := range maps {
			m := maps[j]
			if chk >= m.source && chk <= m.source+m.length-1 {
				chk = m.dest - m.source + chk
				break
			}
		}
	}

	return chk
}

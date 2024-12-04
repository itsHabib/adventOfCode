package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const input = "2024/day2/input.txt"

func main() {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	n := time.Now()

	var (
		s     = bufio.NewScanner(f)
		safe1 int
		safe2 int
	)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}

		nums := ints(line)

		if isSafe(nums) {
			safe1++
			safe2++
			continue
		}

		for i := range nums {
			if isSafe(append(append([]int{}, nums[:i]...), nums[i+1:]...)) {
				safe2++
				break
			}
		}
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("part 1 answer:", safe1)
	fmt.Println("part 2 answer:", safe2)
	fmt.Printf("took %dns\n", time.Since(n).Nanoseconds())
}

func ints(line string) []int {
	parts := strings.Fields(line)
	n := make([]int, len(parts))
	for i := range parts {
		v, err := strconv.Atoi(parts[i])
		if err != nil {
			log.Fatal(err)
		}
		n[i] = v
	}

	return n
}

func isSafe(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	incr := nums[0] < nums[1]
	for i := 0; i < len(nums)-1; i++ {
		if !validLevel(nums[i], nums[i+1], incr) {
			return false
		}
	}

	return true
}

func validLevel(a, b int, incr bool) bool {
	n := a - b
	av := abs(n)

	return (av >= 1 && av <= 3) && ((incr && n < 0) || (!incr && n > 0))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

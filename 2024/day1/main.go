package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
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
		left   = new(queue)
		right  = new(queue)
		rcount = make(map[int]int)
	)
	heap.Init(left)
	heap.Init(right)
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

		heap.Push(left, l)
		heap.Push(right, r)
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	var (
		part1 int
		part2 int
	)
	for left.Len() > 0 && right.Len() > 0 {
		l := heap.Pop(left).(int)
		r := heap.Pop(right).(int)

		part1 += int(math.Abs(float64(r - l)))
		part2 += l * rcount[l]
	}

	fmt.Println("part 1 answer:", part1)
	fmt.Println("part 2 answer:", part2)
}

var _ heap.Interface = (*queue)(nil)

type queue struct {
	items []int
}

func (q queue) Len() int {
	return len(q.items)
}

func (q queue) Less(i, j int) bool {
	return q.items[i] < q.items[j]
}

func (q *queue) Swap(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
}

func (q *queue) Push(x any) {
	q.items = append(q.items, x.(int))
}

func (q *queue) Pop() any {
	x := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return x
}

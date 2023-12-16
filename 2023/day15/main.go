package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("2023/day15/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(scanCommas)
	var ans int
	for scanner.Scan() {
		step := scanner.Text()
		h := hash(step)
		ans += h
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner err:%v", err)
	}

	fmt.Println("part 1 answer:", ans)

	if _, err := f.Seek(0, 0); err != nil {
		log.Fatalf("seek err: %v", err)
	}

	type boxInfo struct {
		label string
		lens  int
	}
	boxes := make(map[int][]boxInfo)
	scanner = bufio.NewScanner(f)
	scanner.Split(scanCommas)
	for scanner.Scan() {
		step := scanner.Text()
		var label string
		var i int
		for i < len(step) {
			if step[i] == '-' || step[i] == '=' {
				label = step[:i]
				break
			}
			i++
		}
		op := step[i]

		box := hash(label)
		switch op {
		case '-':
			info := boxes[box]
			idx := -1
			for i := 0; i < len(info); i++ {
				if info[i].label == label {
					idx = i
					break
				}
			}
			if idx == -1 {
				continue
			}
			info = append(info[:idx], info[idx+1:]...)
			boxes[box] = info
			if len(info) == 0 {
				delete(boxes, box)
			}
		case '=':
			lens := int(step[i+1] - '0')
			info := boxes[box]
			var replaced bool
			for i := range info {
				if info[i].label == label {
					info[i].lens = lens
					replaced = true
					break
				}
			}
			if !replaced {
				info = append(info, boxInfo{label, lens})
			}
			boxes[box] = info
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner err:%v", err)
	}

	var pt2Ans int
	for box, info := range boxes {
		for i := range info {
			pt2Ans += (box + 1) * (i + 1) * info[i].lens
		}
	}

	fmt.Println("part 2 answer:", pt2Ans)
}

func scanCommas(data []byte, atEOF bool) (int, []byte, error) {
	if len(data) == 0 && atEOF {
		return 0, nil, nil
	}

	if i := bytes.IndexByte(data, ','); i >= 0 {
		return i + 1, data[:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}

func hash(s string) int {
	var h int
	for i := range s {
		h += int(s[i])
		h *= 17
		h %= 256
	}

	return h
}

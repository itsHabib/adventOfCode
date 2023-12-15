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
	f, err := os.Open("2023/day12/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var p1Ans int
	var p2Ans int
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		conditionRecord := parts[0]
		numStr := strings.Split(parts[1], ",")
		var nums []int
		for i := range numStr {
			n, _ := strconv.Atoi(numStr[i])
			nums = append(nums, n)
		}
		p1Memo := make(map[[3]int]int)
		p1Ans += getArrangements(conditionRecord, nums, 0, 0, 0, p1Memo)

		var p2Nums []int
		p2NumPart := parts[1]
		for i := 0; i < 4; i++ {
			p2NumPart = p2NumPart + "," + parts[1]
		}
		for _, c := range strings.Split(p2NumPart, ",") {
			n, _ := strconv.Atoi(c)
			p2Nums = append(p2Nums, n)
		}
		p2Memo := make(map[[3]int]int)
		p2Record := conditionRecord
		for i := 0; i < 4; i++ {
			p2Record = p2Record + "?" + conditionRecord
		}
		p2Ans += getArrangements(p2Record, p2Nums, 0, 0, 0, p2Memo)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error during scan: %v", err)
	}

	fmt.Println("part 1 answer:", p1Ans)
	fmt.Println("part 2 answer:", p2Ans)
}

func getArrangements(s string, groups []int, idx, groupIdx, groupSize int, memo map[[3]int]int) int {
	if ansHere, ok := memo[[3]int{idx, groupIdx, groupSize}]; ok {
		return ansHere
	}
	if idx == len(s) {
		if (groupIdx == len(groups) && groupSize == 0) || (groupIdx == len(groups)-1 && groups[groupIdx] == groupSize) {
			return 1
		}
		return 0
	}

	var ans int
	if s[idx] == '.' || s[idx] == '?' {
		if groupSize == 0 {
			ans += getArrangements(s, groups, idx+1, groupIdx, 0, memo)
		} else if groupSize > 0 && groupIdx < len(groups) && groups[groupIdx] == groupSize {
			ans += getArrangements(s, groups, idx+1, groupIdx+1, 0, memo)
		}
	}
	if s[idx] == '#' || s[idx] == '?' {
		ans += getArrangements(s, groups, idx+1, groupIdx, groupSize+1, memo)
	}

	memo[[3]int{idx, groupIdx, groupSize}] = ans

	return ans
}

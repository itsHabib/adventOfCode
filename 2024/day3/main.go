package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const input = "2024/day3/input.txt"

var disabled = false

// mul(x,y)

func main() {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}

	var (
		s    = bufio.NewScanner(f)
		sum1 int
		sum2 int
		p1   = new(parser)
		p2   = new(parser)
	)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}

		p1.reset(line)
		p2.reset(line)

		sum1 += p1.parse(false)
		sum2 += p2.parse(true)
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("part 1 answer:", sum1)
	fmt.Println("part 2 answer:", sum2)
}

type parser struct {
	s   string
	cur int
}

func (p *parser) reset(s string) {
	p.s = s
	p.cur = 0
}

func (p *parser) parse(withDos bool) int {
	var sum int

	for p.cur < len(p.s) {
		switch {
		case p.s[p.cur] == 'm':
			n, ok := p.mul()
			if ok && (!disabled || !withDos) {
				sum += n
			}
		case p.s[p.cur] == 'd' && withDos:
			p.do()
		default:
			p.next()
		}
	}

	return sum
}

func (p *parser) mul() (int, bool) {
	if !strings.HasPrefix(p.s[p.cur:], "mul") {
		p.next()
		return 0, false
	}

	p.jump(len("mul"))

	if !p.at('(') {
		return 0, false
	}

	p.next()

	x, ok := p.number()
	if !ok {
		return 0, false
	}

	if !p.at(',') {
		return 0, false
	}

	p.next()

	y, ok := p.number()
	if !ok {
		return 0, false
	}

	if !p.at(')') {
		return 0, false
	}

	p.next()

	return x * y, true
}

func (p *parser) do() {
	if !strings.HasPrefix(p.s[p.cur:], "do") {
		p.next()
		return
	}

	p.jump(len("do"))

	switch {
	case p.scan("()"):
		p.jump(2)
		disabled = false
	case p.scan("n't()"):
		p.jump(5)
		disabled = true
	}
}

func (p *parser) number() (int, bool) {
	ns := ""

	for p.cur < len(p.s) {
		if p.s[p.cur] == ',' || p.s[p.cur] == ')' {
			break
		}

		if p.s[p.cur] < '0' || p.s[p.cur] > '9' {
			return 0, false
		}

		ns += string(p.s[p.cur])

		p.next()
	}

	n, err := strconv.Atoi(ns)
	if err != nil {
		return 0, false
	}

	return n, true
}

func (p *parser) at(ch byte) bool {
	return p.cur >= 0 && p.cur < len(p.s) && p.s[p.cur] == ch
}

func (p *parser) scan(sub string) bool {
	return p.cur < len(p.s) && strings.HasPrefix(p.s[p.cur:p.cur+len(sub)], sub)
}

func (p *parser) next() {
	p.cur++
}

func (p *parser) jump(n int) {
	p.cur += n
}

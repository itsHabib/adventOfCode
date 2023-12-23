package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type workflow struct {
	name  string
	rules []rule
}

func (wf workflow) run(param workflowParam) string {
	for i := range wf.rules {
		r := wf.rules[i]
		if r.op == 0 {
			return r.outcome
		}
		var rating int
		switch r.rating {
		case "x":
			rating = param.x
		case "m":
			rating = param.m
		case "a":
			rating = param.a
		case "s":
			rating = param.s
		}
		var isTrue bool
		switch r.op {
		case '>':
			isTrue = rating > r.n
		case '<':
			isTrue = rating < r.n
		}
		if isTrue {
			return r.outcome
		}
	}
	return "R"
}

type workflowParam struct {
	x, m, a, s int
}
type rule struct {
	rating  string
	op      byte
	n       int
	outcome string
}

type workflowRange struct {
	min int
	max int
}

func main() {
	f, err := os.Open("2023/day19/input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	workflows := make(map[string]workflow)
	for sc.Scan() {
		if sc.Text() == "" {
			break
		}
		wf := parseWorkflow(sc.Text())
		workflows[wf.name] = wf
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scanner err:%v", err)
	}
	var params []workflowParam
	for sc.Scan() {
		paramParts := strings.Split(strings.TrimFunc(sc.Text(), func(r rune) bool {
			return r == '{' || r == '}'
		}), ",")
		var p workflowParam
		for i := range paramParts {
			ratingParts := strings.Split(paramParts[i], "=")
			switch ratingParts[0] {
			case "x":
				p.x, _ = strconv.Atoi(ratingParts[1])
			case "m":
				p.m, _ = strconv.Atoi(ratingParts[1])
			case "a":
				p.a, _ = strconv.Atoi(ratingParts[1])
			case "s":
				p.s, _ = strconv.Atoi(ratingParts[1])
			}
		}
		params = append(params, p)
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scanner err: %v", err)
	}

	fmt.Println("part 1 answer:", part1(workflows, params))
	fmt.Println("part 2 answer:", runWorkflowsPt2(workflows))
}

func part1(workflows map[string]workflow, params []workflowParam) int {
	var sum int
	for i := range params {
		accepted := runWorkflows(workflows, params[i])
		if accepted {
			sum += params[i].x + params[i].m + params[i].a + params[i].s
		}
	}

	return sum
}

func runWorkflows(workflows map[string]workflow, p workflowParam) bool {
	run := []workflow{workflows["in"]}

	for len(run) > 0 {
		wf := run[len(run)-1]
		run = run[:len(run)-1]
		outcome := wf.run(p)
		switch outcome {
		case "A":
			return true
		case "R":
			return false
		default:
			run = append(run, workflows[outcome])
		}
	}

	return false
}

type runRange struct {
	name string
	x    workflowRange
	m    workflowRange
	a    workflowRange
	s    workflowRange
}

func runWorkflowsPt2(workflows map[string]workflow) int {
	run := []runRange{{
		name: "in",
		x:    workflowRange{1, 4000},
		m:    workflowRange{1, 4000},
		a:    workflowRange{1, 4000},
		s:    workflowRange{1, 4000},
	}}

	var combos int
	for len(run) > 0 {
		rr := run[0]
		run = run[1:]
		if rr.x.min > rr.x.max || rr.m.min > rr.m.max || rr.a.min > rr.a.max || rr.s.min > rr.s.max {
			continue
		}
		switch rr.name {
		case "A":
			combos += (rr.x.max - rr.x.min + 1) * (rr.m.max - rr.m.min + 1) * (rr.a.max - rr.a.min + 1) * (rr.s.max - rr.s.min + 1)
		case "R":
			continue
		default:
			wf := workflows[rr.name]
			for i := range wf.rules {
				r := wf.rules[i]
				if r.op == 0 {
					run = append(run, runRange{
						name: r.outcome,
						x:    rr.x,
						m:    rr.m,
						a:    rr.a,
						s:    rr.s,
					})
					break
				}
				x, m, a, s := newRanges(r.rating, string(r.op), r.n, rr.x, rr.m, rr.a, rr.s)
				run = append(run, runRange{
					name: r.outcome,
					x:    x,
					m:    m,
					a:    a,
					s:    s,
				})
				switch r.op {
				case '>':
					rr.x, rr.m, rr.a, rr.s = newRanges(r.rating, "<=", r.n, rr.x, rr.m, rr.a, rr.s)
				case '<':
					rr.x, rr.m, rr.a, rr.s = newRanges(r.rating, ">=", r.n, rr.x, rr.m, rr.a, rr.s)
				}
			}
		}
	}

	return combos
}

func parseWorkflow(s string) workflow {
	firstBracket := strings.Index(s, "{")
	name := s[:firstBracket]
	rules := strings.Split(strings.TrimFunc(s[firstBracket:], func(r rune) bool {
		return r == '{' || r == '}'
	}), ",")
	var wfRules []rule
	for i := range rules {
		wfRules = append(wfRules, parseRule(rules[i]))
	}

	return workflow{
		name:  name,
		rules: wfRules,
	}
}

func parseRule(s string) rule {
	var (
		start int
		cur   = start
	)
	for cur < len(s) && s[cur] != '<' && s[cur] != '>' {
		cur++
	}
	if cur == len(s) {
		return rule{
			outcome: s,
		}
	}

	rating := s[start:cur]
	op := s[cur]

	cur++
	start = cur
	for s[cur] != ':' {
		cur++
	}
	n, _ := strconv.Atoi(s[start:cur])

	return rule{
		rating:  rating,
		op:      op,
		n:       n,
		outcome: s[cur+1:],
	}
}

func newRange(op string, n int, r workflowRange) workflowRange {
	switch op {
	case ">":
		r.min = max(r.min, n+1)
	case "<":
		r.max = min(r.max, n-1)
	case ">=":
		r.min = max(r.min, n)
	case "<=":
		r.max = min(r.max, n)
	}
	return r
}

func newRanges(rating string, op string, n int, x, m, a, s workflowRange) (workflowRange, workflowRange, workflowRange, workflowRange) {
	switch rating {
	case "x":
		x = newRange(op, n, x)
	case "m":
		m = newRange(op, n, m)
	case "a":
		a = newRange(op, n, a)
	case "s":
		s = newRange(op, n, s)
	}

	return x, m, a, s
}

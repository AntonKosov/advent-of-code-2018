package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2018/aoc/heap"
	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
)

const (
	stepTime   = 60
	maxWorkers = 5
	// stepTime   = 0
	// maxWorkers = 2
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (instructions map[rune][]rune) {
	lines := input.Lines()
	lines = lines[:len(lines)-1]
	instructions = make(map[rune][]rune, 26)
	for _, line := range lines {
		first, second := rune(line[5]), rune(line[36])
		instructions[first] = append(instructions[first], second)
	}

	return
}

func process(instructions map[rune][]rune) int {
	rank := make(map[rune]int, len(instructions))
	for _, steps := range instructions {
		for _, r := range steps {
			rank[r]++
		}
	}

	pq := heap.New(func(r1, r2 rune) bool { return r1 < r2 })
	for step := range instructions {
		if rank[step] == 0 {
			pq.Push(step)
		}
	}

	totalTime := 0
	workers := heap.New(func(w1, w2 Worker) bool { return w1.endTime < w2.endTime })
	for !pq.Empty() || !workers.Empty() {
		for !pq.Empty() && workers.Len() < maxWorkers {
			step := pq.Pop()
			workers.Push(Worker{step: step, endTime: totalTime + stepTime + int(step-'A') + 1})
		}
		for {
			worker := workers.Pop()
			totalTime = worker.endTime
			step := worker.step
			steps := instructions[step]
			for _, step := range steps {
				rank[step]--
				if rank[step] == 0 {
					pq.Push(step)
				}
			}
			if workers.Empty() || workers.Peek().endTime != worker.endTime {
				break
			}
		}
	}

	return totalTime
}

type Worker struct {
	endTime int
	step    rune
}

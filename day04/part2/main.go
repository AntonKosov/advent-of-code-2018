package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() map[int][]SleepTime {
	lines := input.Lines()
	lines = lines[:len(lines)-1]
	slices.Sort(lines)
	guards := map[int][]SleepTime{}
	for len(lines) > 0 {
		guardID := transform.StrToInts(lines[0][26:])[0]
		lines = lines[1:]
		for len(lines) > 0 && strings.Contains(lines[0], "falls") {
			guards[guardID] = append(guards[guardID], SleepTime{
				fallsAsleep: transform.StrToInts(lines[0][15:17])[0],
				wakeUp:      transform.StrToInts(lines[1][15:17])[0],
			})
			lines = lines[2:]
		}
	}

	return guards
}

func process(guards map[int][]SleepTime) int {
	bestResult := 0
	bestTimes := 0
	for id, sleepTimes := range guards {
		minutes := [60]int{}
		for _, sleepTime := range sleepTimes {
			for i := sleepTime.fallsAsleep; i < sleepTime.wakeUp; i++ {
				minutes[i]++
			}
		}
		idx, times := maxTimes(minutes)
		if bestTimes < times {
			bestTimes = times
			bestResult = id * idx
		}
	}

	return bestResult
}

func maxTimes(minutes [60]int) (idx, times int) {
	for i := 1; i < len(minutes); i++ {
		value := minutes[i]
		if value > times {
			times = value
			idx = i
		}
	}

	return
}

type SleepTime struct {
	fallsAsleep int
	wakeUp      int
}

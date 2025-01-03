package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
)

const (
	generations               = 50_000_000_000
	maxStateLength            = 1_000
	patternDetectionThreshold = 100
	plant                     = '#'
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (state []uint8, rules map[uint8]uint8) {
	stringToBits := func(str string) []uint8 {
		bits := make([]uint8, len(str))
		for i, v := range str {
			var value uint8
			if v == plant {
				value = 1
			}
			bits[i] = value
		}
		return bits
	}

	lines := input.Lines()
	state = stringToBits(lines[0][15:])
	lines = lines[2 : len(lines)-1]
	rules = make(map[uint8]uint8, len(lines))
	for _, line := range lines {
		if line[9] == plant {
			rules[stateToKey(stringToBits(line[:5]))] = 1
		}
	}

	return
}

func process(state []uint8, rules map[uint8]uint8) int64 {
	negative := 0
	var scores []int64
	var patternValue int64
	patternIdx := 0
	scores = append(scores, calcSum(state, 0))
	for len(state) < maxStateLength {
		if stateSum(state[:4]) > 0 {
			state = append([]uint8{0, 0, 0, 0}, state...)
			negative += 4
		}
		if stateSum(state[len(state)-4:]) > 0 {
			state = append(state, 0, 0, 0, 0)
		}
		nextState := make([]uint8, 2, len(state))
		for i, key := 2, stateToKey(state[:5]); i < len(state)-2; i++ {
			nextState = append(nextState, rules[key])
			if i < len(state)-3 {
				key = ((key & 0b1111) << 1) | state[i+3]
			}
		}
		state = append(nextState, 0, 0)
		sum := calcSum(state, negative)
		if diff := sum - scores[len(scores)-1]; diff == int64(patternValue) {
			if len(scores)-patternIdx == patternDetectionThreshold {
				return scores[patternIdx-1] + patternValue*(generations+1-int64(patternIdx))
			}
		} else {
			patternValue = diff
			patternIdx = len(scores)
		}
		scores = append(scores, sum)
	}

	panic("pattern not found")
}

func calcSum(state []uint8, negative int) (sum int64) {
	for i, v := range state {
		if v == 1 {
			sum += int64(i - negative)
		}
	}

	return
}

func stateSum(state []uint8) (result int) {
	for _, v := range state {
		result += int(v)
	}

	return
}

func stateToKey(state []uint8) (key uint8) {
	for i, s := range state {
		key |= s << (len(state) - i - 1)
	}

	return
}

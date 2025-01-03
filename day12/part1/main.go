package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
)

const (
	generations = 20
	plant       = '#'
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

func process(state []uint8, rules map[uint8]uint8) int {
	negative := 0
	for i := 0; i < generations; i++ {
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
	}

	sum := 0
	for i, v := range state {
		if v == 1 {
			sum += i - negative
		}
	}

	return sum
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

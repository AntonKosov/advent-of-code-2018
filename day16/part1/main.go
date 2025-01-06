package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []Sample {
	lines := input.Lines()
	var samples []Sample
	for len(lines) > 0 {
		sampleLines := lines[:4]
		if len(sampleLines[0]) == 0 {
			break
		}
		lines = lines[len(sampleLines):]
		before := transform.StrToInts(sampleLines[0])
		instruction := transform.StrToInts(sampleLines[1])
		after := transform.StrToInts(sampleLines[2])
		samples = append(samples, Sample{
			registersBefore: Registers{before[0], before[1], before[2], before[3]},
			instruction: Instruction{
				opcode: instruction[0],
				a:      instruction[1],
				b:      instruction[2],
				c:      instruction[3],
			},
			registersAfter: Registers{after[0], after[1], after[2], after[3]},
		})
	}

	return samples
}

func process(samples []Sample) int {
	count := 0
	for _, sample := range samples {
		if countOptions(sample) >= 3 {
			count++
		}
	}

	return count
}

func countOptions(sample Sample) int {
	count := 0
	for _, op := range Operations {
		registers := sample.registersBefore
		op(sample.instruction, &registers)
		if registers == sample.registersAfter {
			count++
		}
	}

	return count
}

type Operation func(Instruction, *Registers)

var Operations = []Operation{
	// addr
	func(i Instruction, r *Registers) { r[i.c] = r[i.a] + r[i.b] },
	// addi
	func(i Instruction, r *Registers) { r[i.c] = r[i.a] + i.b },
	// mulr
	func(i Instruction, r *Registers) { r[i.c] = r[i.a] * r[i.b] },
	// muli
	func(i Instruction, r *Registers) { r[i.c] = r[i.a] * i.b },
	// banr
	func(i Instruction, r *Registers) { r[i.c] = r[i.a] & r[i.b] },
	// bani
	func(i Instruction, r *Registers) { r[i.c] = r[i.a] & i.b },
	// borr
	func(i Instruction, r *Registers) { r[i.c] = r[i.a] | r[i.b] },
	// bori
	func(i Instruction, r *Registers) { r[i.c] = r[i.a] | i.b },
	// setr
	func(i Instruction, r *Registers) { r[i.c] = r[i.a] },
	// seti
	func(i Instruction, r *Registers) { r[i.c] = i.a },
	// gtir
	func(i Instruction, r *Registers) { r[i.c] = cmp(i.a > r[i.b]) },
	// gtri
	func(i Instruction, r *Registers) { r[i.c] = cmp(r[i.a] > i.b) },
	// gtrr
	func(i Instruction, r *Registers) { r[i.c] = cmp(r[i.a] > r[i.b]) },
	// eqir
	func(i Instruction, r *Registers) { r[i.c] = cmp(i.a == r[i.b]) },
	// eqri
	func(i Instruction, r *Registers) { r[i.c] = cmp(r[i.a] == i.b) },
	// eqrr
	func(i Instruction, r *Registers) { r[i.c] = cmp(r[i.a] == r[i.b]) },
}

func cmp(cond bool) int {
	if cond {
		return 1
	}

	return 0
}

type Registers [4]int

type Instruction struct {
	opcode int
	a      int
	b      int
	c      int
}

type Sample struct {
	registersBefore Registers
	registersAfter  Registers
	instruction     Instruction
}

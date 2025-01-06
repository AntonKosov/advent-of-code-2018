package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/maps"
	"github.com/AntonKosov/advent-of-code-2018/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2018/aoc/stack"
	"github.com/AntonKosov/advent-of-code-2018/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (samples []Sample, program []Instruction) {
	lines := input.Lines()
	lines = lines[:len(lines)-1]
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

	program = slice.Map(lines[2:], func(line string) Instruction {
		parts := transform.StrToInts(line)
		return Instruction{
			opcode: parts[0],
			a:      parts[1],
			b:      parts[2],
			c:      parts[3],
		}
	})

	return
}

func process(samples []Sample, program []Instruction) int {
	opcodes := detectOpcodes(samples)
	var registers Registers
	for _, line := range program {
		Operations[opcodes[line.opcode]](line, &registers)
	}

	return registers[0]
}

func detectOpcodes(samples []Sample) []int {
	opcodes := make([]map[int]bool, 16)
	result := make([]int, len(opcodes))
	for i := range opcodes {
		m := make(map[int]bool, len(Operations))
		for i := range Operations {
			m[i] = true
		}
		opcodes[i] = m
	}

	for _, sample := range samples {
		opcode := sample.instruction.opcode
		possibleOperations := opcodes[opcode]
		if len(possibleOperations) == 1 {
			continue
		}
		testOptions(sample, possibleOperations)
		if len(possibleOperations) > 1 {
			continue
		}
		foundOperation := maps.SingleKey(possibleOperations)
		result[opcode] = foundOperation
		foundOperations := stack.New[int]()
		foundOperations.Push(foundOperation)
		for !foundOperations.Empty() {
			fo := foundOperations.Pop()
			for opcode, possibleOperations := range opcodes {
				if len(possibleOperations) == 1 {
					continue
				}
				delete(possibleOperations, fo)
				if len(possibleOperations) == 1 {
					fo2 := maps.SingleKey(possibleOperations)
					result[opcode] = fo2
					foundOperations.Push(fo2)
				}
			}
		}
	}

	return result
}

func testOptions(sample Sample, possibleOperationIndeces map[int]bool) {
	for i, op := range Operations {
		if !possibleOperationIndeces[i] {
			continue
		}
		registers := sample.registersBefore
		op(sample.instruction, &registers)
		if registers != sample.registersAfter {
			delete(possibleOperationIndeces, i)
		}
	}
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

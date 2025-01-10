package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2018/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (instructionPointerRegister int, instructions []Instruction) {
	lines := input.Lines()
	instructionPointerRegister = transform.StrToInts(lines[0])[0]
	lines = lines[1 : len(lines)-1]
	instructions = slice.Map(lines, func(line string) Instruction {
		parts := strings.Split(line, " ")
		return Instruction{
			opcode: parts[0],
			a:      transform.StrToInt(parts[1]),
			b:      transform.StrToInt(parts[2]),
			c:      transform.StrToInt(parts[3]),
		}
	})

	return
}

func process(instructionPointerRegister int, instructions []Instruction) int {
	var registers Registers
	for registers[instructionPointerRegister] < len(instructions) {
		instruction := instructions[registers[instructionPointerRegister]]
		Operations[instruction.opcode](instruction, &registers)
		registers[instructionPointerRegister]++
	}

	return registers[0]
}

type Registers [6]int

type Instruction struct {
	opcode string
	a      int
	b      int
	c      int
}

func cmp(cond bool) int {
	if cond {
		return 1
	}

	return 0
}

type Operation func(Instruction, *Registers)

var Operations = map[string]Operation{
	"addr": func(i Instruction, r *Registers) { r[i.c] = r[i.a] + r[i.b] },
	"addi": func(i Instruction, r *Registers) { r[i.c] = r[i.a] + i.b },
	"mulr": func(i Instruction, r *Registers) { r[i.c] = r[i.a] * r[i.b] },
	"muli": func(i Instruction, r *Registers) { r[i.c] = r[i.a] * i.b },
	"banr": func(i Instruction, r *Registers) { r[i.c] = r[i.a] & r[i.b] },
	"bani": func(i Instruction, r *Registers) { r[i.c] = r[i.a] & i.b },
	"borr": func(i Instruction, r *Registers) { r[i.c] = r[i.a] | r[i.b] },
	"bori": func(i Instruction, r *Registers) { r[i.c] = r[i.a] | i.b },
	"setr": func(i Instruction, r *Registers) { r[i.c] = r[i.a] },
	"seti": func(i Instruction, r *Registers) { r[i.c] = i.a },
	"gtir": func(i Instruction, r *Registers) { r[i.c] = cmp(i.a > r[i.b]) },
	"gtri": func(i Instruction, r *Registers) { r[i.c] = cmp(r[i.a] > i.b) },
	"gtrr": func(i Instruction, r *Registers) { r[i.c] = cmp(r[i.a] > r[i.b]) },
	"eqir": func(i Instruction, r *Registers) { r[i.c] = cmp(i.a == r[i.b]) },
	"eqri": func(i Instruction, r *Registers) { r[i.c] = cmp(r[i.a] == i.b) },
	"eqrr": func(i Instruction, r *Registers) { r[i.c] = cmp(r[i.a] == r[i.b]) },
}

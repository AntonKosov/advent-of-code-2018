package main

import (
	"fmt"
	"slices"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/math"
	"github.com/AntonKosov/advent-of-code-2018/aoc/slice"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (paths [][]rune, carts []Cart) {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	dirs := map[rune]math.Vector2[int]{'<': leftDir, '>': rightDir, '^': upDir, 'v': downDir}
	pathRecover := map[rune]rune{'<': '-', '>': '-', '^': '|', 'v': '|'}
	paths = slice.New2D[rune](len(lines), len(lines[0]))
	for r, line := range lines {
		for c, v := range line {
			if dir, ok := dirs[v]; ok {
				carts = append(carts, Cart{
					position:  math.NewVector2(c, r),
					direction: dir,
				})
				v = pathRecover[v]
			}
			paths[r][c] = v
		}
	}

	return
}

func process(paths [][]rune, carts []Cart) string {
	positions := make(map[math.Vector2[int]]bool, len(carts))
	for _, cart := range carts {
		positions[cart.position] = true
	}

	for {
		sortCarts(carts)
		for i := range carts {
			cart := &carts[i]
			delete(positions, cart.position)
			moveCart(paths, cart)
			if pos := cart.position; positions[pos] {
				return fmt.Sprintf("%v,%v", pos.X, pos.Y)
			}
			positions[cart.position] = true
		}
	}
}

func moveCart(paths [][]rune, cart *Cart) {
	pos := cart.position.Add(cart.direction)
	cart.position = pos
	var rotation Rotation
	cell := paths[pos.Y][pos.X]
	switch cell {
	case '/':
		fallthrough
	case '\\':
		rotation = curves[Curve{cell: cell, entryDir: cart.direction}]
	case '+':
		rotation = cart.nextRotation
		cart.nextRotation = (cart.nextRotation + 1) % 3
	default:
		return
	}

	switch rotation {
	case LeftRotation:
		cart.direction = cart.direction.RotateLeft()
	case RightRotation:
		cart.direction = cart.direction.RotateRight()
	}
}

func sortCarts(carts []Cart) {
	slices.SortFunc(carts, func(c1, c2 Cart) int {
		if c1.position.Y != c2.position.Y {
			return c1.position.Y - c2.position.Y
		}

		return c1.position.X - c2.position.X
	})
}

type Rotation int

const (
	LeftRotation  Rotation = iota
	NoRotation    Rotation = iota
	RightRotation Rotation = iota
)

type Cart struct {
	position     math.Vector2[int]
	direction    math.Vector2[int]
	nextRotation Rotation
}

type Curve struct {
	cell     rune
	entryDir math.Vector2[int]
}

var (
	leftDir  = math.NewVector2(-1, 0)
	rightDir = math.NewVector2(1, 0)
	upDir    = math.NewVector2(0, -1)
	downDir  = math.NewVector2(0, 1)
	curves   map[Curve]Rotation
)

func init() {
	curves = map[Curve]Rotation{
		{cell: '\\', entryDir: leftDir}:  RightRotation,
		{cell: '\\', entryDir: rightDir}: RightRotation,
		{cell: '\\', entryDir: upDir}:    LeftRotation,
		{cell: '\\', entryDir: downDir}:  LeftRotation,
		{cell: '/', entryDir: leftDir}:   LeftRotation,
		{cell: '/', entryDir: rightDir}:  LeftRotation,
		{cell: '/', entryDir: upDir}:     RightRotation,
		{cell: '/', entryDir: downDir}:   RightRotation,
	}
}

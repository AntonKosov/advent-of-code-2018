package main

import (
	"fmt"
	smath "math"
	"strings"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/math"
	"github.com/AntonKosov/advent-of-code-2018/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2018/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("%v\n", answer)
}

func read() []Point {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	return slice.Map(lines, func(line string) Point {
		parts := transform.StrToInts(line)
		return Point{
			position: math.NewVector2(parts[0], parts[1]),
			velocity: math.NewVector2(parts[2], parts[3]),
		}
	})
}

func process(points []Point) string {
	for size := smath.MaxInt; ; {
		move(points, 1)
		leftTop, bottomRight := bounds(points)
		size2 := leftTop.ManhattanDst(bottomRight)
		if size2 > size {
			break
		}
		size = size2
	}

	move(points, -1)

	return pointsToImage(points)
}

func pointsToImage(points []Point) string {
	leftTop, bottomRight := bounds(points)
	height := bottomRight.Y - leftTop.Y + 1
	width := bottomRight.X - leftTop.X + 1
	img := slice.New2D[rune](width, height)
	slice.Fill2D(img, ' ')
	for _, p := range points {
		pos := p.position.Sub(leftTop)
		img[pos.Y][pos.X] = '#'
	}

	rows := make([]string, height)
	for i := range height {
		rows[i] = string(img[i])
	}

	return strings.Join(rows, "\n")
}

func bounds(points []Point) (leftTop, bottomRight math.Vector2[int]) {
	leftTop = points[0].position
	bottomRight = leftTop
	for _, point := range points {
		pos := point.position
		leftTop.X = min(leftTop.X, pos.X)
		leftTop.Y = min(leftTop.Y, pos.Y)
		bottomRight.X = max(bottomRight.X, pos.X)
		bottomRight.Y = max(bottomRight.Y, pos.Y)
	}

	return
}

func move(points []Point, steps int) {
	for i := range points {
		points[i].Move(steps)
	}
}

type Point struct {
	position, velocity math.Vector2[int]
}

func (p *Point) Move(steps int) {
	p.position = p.position.Add(p.velocity.Mul(steps))
}

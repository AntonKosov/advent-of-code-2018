package main

import (
	"fmt"
	smath "math"
	"strings"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/math"
	"github.com/AntonKosov/advent-of-code-2018/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (tiles Tiles, maxY int) {
	lines := input.Lines()
	lines = lines[:len(lines)-1]
	tiles = Tiles{}
	for _, line := range lines {
		nums := transform.StrToInts(line)
		isX := line[0] == 'x'
		for i := nums[1]; i <= nums[2]; i++ {
			var pos math.Vector2[int]
			if isX {
				pos = math.NewVector2(nums[0], i)
			} else {
				pos = math.NewVector2(i, nums[0])
			}
			tiles[pos] = WallTile
			maxY = max(maxY, pos.Y)
		}
	}

	return
}

func process(tiles Tiles, maxY int) int {
	count := countWaterTiles(tiles, math.NewVector2(500, 1), maxY)
	fmt.Println(tiles)

	return count
}

func countWaterTiles(tiles Tiles, pos math.Vector2[int], maxY int) int {
	if pos.Y > maxY || tiles[pos] != EmptyTile {
		return 0
	}

	tiles[pos] = RunningWaterTile
	below := pos.Add(Down)
	count := countWaterTiles(tiles, below, maxY)
	if tiles[below] == RunningWaterTile || pos.Y == maxY {
		return count
	}

	spread := func(dir math.Vector2[int]) (x int, rest bool) {
		for pos := pos; ; {
			nextPos := pos.Add(dir)
			if tiles[nextPos] != EmptyTile {
				return pos.X, true
			}
			tiles[nextPos] = RunningWaterTile
			below := nextPos.Add(Down)
			count += countWaterTiles(tiles, below, maxY)
			if tiles[below] == RunningWaterTile {
				return nextPos.X, false
			}
			pos = nextPos
		}
	}

	leftX, restLeft := spread(Left)
	rightX, restRight := spread(Right)
	if restLeft && restRight {
		for x := leftX; x <= rightX; x++ {
			tiles[math.NewVector2(x, pos.Y)] = RestWaterTile
		}
		count += rightX - leftX + 1
	}

	return count
}

type Tile uint8

type Tiles map[math.Vector2[int]]Tile

func (t Tiles) String() string {
	minX, maxX, minY, maxY := smath.MaxInt, 0, smath.MaxInt, 0
	for pos := range t {
		minX = min(minX, pos.X)
		maxX = max(maxX, pos.X)
		minY = min(minY, pos.Y)
		maxY = max(maxY, pos.Y)
	}

	letters := map[Tile]rune{
		EmptyTile:        '.',
		WallTile:         '#',
		RunningWaterTile: '|',
		RestWaterTile:    '~',
	}
	var sb strings.Builder
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			sb.WriteRune(letters[t[math.NewVector2(x, y)]])
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

const (
	EmptyTile        Tile = iota
	WallTile         Tile = iota
	RunningWaterTile Tile = iota
	RestWaterTile    Tile = iota
)

var (
	Down  = math.NewVector2(0, 1)
	Left  = math.NewVector2(-1, 0)
	Right = math.NewVector2(1, 0)
)

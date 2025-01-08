package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/slice"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() [][]Tile {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	return slice.Map(lines, func(line string) []Tile { return []Tile(line) })
}

func process(tiles [][]Tile) int {
	printTiles(tiles)
	nextTiles := slice.New2D[Tile](len(tiles), len(tiles[0]))
	for i := 0; i < 10; i++ {
		for r, row := range tiles {
			for c := range row {
				nextTiles[r][c] = transformTile(tiles, r, c)
			}
		}
		tiles, nextTiles = nextTiles, tiles
		printTiles(tiles)
	}

	woodedAcres, lumberyards := 0, 0
	for _, row := range tiles {
		for _, v := range row {
			switch v {
			case Trees:
				woodedAcres++
			case Lumberyard:
				lumberyards++
			}
		}
	}

	return woodedAcres * lumberyards
}

func transformTile(tiles [][]Tile, r, c int) Tile {
	s := make(map[Tile]int, 3)
	for y := max(0, r-1); y <= min(len(tiles)-1, r+1); y++ {
		for x := max(0, c-1); x <= min(len(tiles[y])-1, c+1); x++ {
			if y == r && x == c {
				continue
			}
			s[tiles[y][x]]++
		}
	}

	ct := tiles[r][c]
	switch ct {
	case OpenGround:
		if s[Trees] >= 3 {
			return Trees
		}
	case Trees:
		if s[Lumberyard] >= 3 {
			return Lumberyard
		}
	case Lumberyard:
		if s[Lumberyard] == 0 || s[Trees] == 0 {
			return OpenGround
		}
	}

	return ct
}

type Tile rune

const (
	OpenGround Tile = '.'
	Trees      Tile = '|'
	Lumberyard Tile = '#'
)

func printTiles(tiles [][]Tile) {
	for _, row := range tiles {
		fmt.Println(string(row))
	}
	fmt.Println()
}

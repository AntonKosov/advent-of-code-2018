package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/math"
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
	const seconds = 1_000_000_000
	cachedKeys := []Key{getKey(tiles)}
	cache := map[Key]int{cachedKeys[0]: 0}
	printTiles(tiles)
	nextTiles := slice.New2D[Tile](len(tiles), len(tiles[0]))
	for i := 0; i < seconds; i++ {
		for r, row := range tiles {
			for c := range row {
				nextTiles[r][c] = transformTile(tiles, r, c)
			}
		}
		tiles, nextTiles = nextTiles, tiles
		printTiles(tiles)
		key := getKey(tiles)
		if idx, ok := cache[key]; ok {
			cachedKeys = cachedKeys[idx:]
			targetKey := cachedKeys[(seconds-idx)%len(cachedKeys)]
			trees, lumberyards := countTiles(targetKey)
			return trees * lumberyards
		}
		cachedKeys = append(cachedKeys, key)
		cache[key] = len(cachedKeys) - 1
	}

	panic("loop not found")
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

func countTiles(key Key) (trees, lumberyards int) {
	for _, row := range key {
		trees += math.CountBits(row[0])
		lumberyards += math.CountBits(row[1])
	}

	return
}

func getKey(tiles [][]Tile) (key Key) {
	for r, row := range tiles {
		for c, v := range row {
			switch v {
			case Trees:
				key[r][0] |= 1 << c
			case Lumberyard:
				key[r][1] |= 1 << c
			}
		}
	}

	return
}

type Tile rune

const (
	OpenGround Tile = '.'
	Trees      Tile = '|'
	Lumberyard Tile = '#'
)

type Key [50][2]uint64

func printTiles(tiles [][]Tile) {
	for _, row := range tiles {
		fmt.Println(string(row))
	}
	fmt.Println()
}

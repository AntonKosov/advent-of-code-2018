package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2018/aoc/heap"
	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/math"
	"github.com/AntonKosov/advent-of-code-2018/aoc/transform"
)

const (
	MoveMinutes       = 1
	GearChangeMinutes = 7
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (depth int, target math.Vector2[int]) {
	lines := input.Lines()
	depth = transform.StrToInts(lines[0])[0]
	targetValues := transform.StrToInts(lines[1])
	target = math.NewVector2(targetValues[0], targetValues[1])

	return
}

func process(depth int, target math.Vector2[int]) int {
	cave := NewCave(depth, target)
	states := newHeap(target)
	states.Push(State{Position: math.NewVector2(0, 0), Item: TorchItem})
	visitedStates := map[VisitedState]int{}
	for !states.Empty() {
		state := states.Pop()
		vs := VisitedState{Position: state.Position, Item: state.Item}
		if mins, ok := visitedStates[vs]; ok && mins <= state.Minutes {
			continue
		}
		visitedStates[vs] = state.Minutes

		if state.Position == target {
			if state.Item != TorchItem {
				state.Minutes += GearChangeMinutes
			}

			return state.Minutes
		}
		for _, dir := range Dirs {
			nextPos := state.Position.Add(dir)
			if nextPos.X < 0 || nextPos.Y < 0 {
				continue
			}
			nextItem, minutes := prepareToMove(cave.Region(state.Position), cave.Region(nextPos), state.Item)
			states.Push(State{
				Position: nextPos,
				Item:     nextItem,
				Minutes:  state.Minutes + MoveMinutes + minutes,
			})
		}
	}

	panic("path not found")
}

func prepareToMove(currentRegion, nextRegion Region, currentItem Item) (nextItems Item, extraMinutes int) {
	allowedItemsNextRegion := AllowedItems[nextRegion]
	if currentRegion == nextRegion || allowedItemsNextRegion[currentItem] {
		return currentItem, 0
	}

	allowedItemsCurrentRegion := AllowedItems[currentRegion]
	for item := range allowedItemsNextRegion {
		if allowedItemsCurrentRegion[item] && allowedItemsNextRegion[item] {
			return item, GearChangeMinutes
		}
	}

	panic("item not found")
}

func newHeap(target math.Vector2[int]) *heap.Heap[State] {
	return heap.New(func(a, b State) bool {
		av := a.Position.ManhattanDst(target) + a.Minutes
		if a.Item != TorchItem {
			av += GearChangeMinutes
		}
		bv := b.Position.ManhattanDst(target) + b.Minutes
		if b.Item != TorchItem {
			bv += GearChangeMinutes
		}
		return av < bv
	})
}

type VisitedState struct {
	Position math.Vector2[int]
	Item     Item
}

type State struct {
	Position math.Vector2[int]
	Item     Item
	Minutes  int
}

type Region uint8

const (
	RockyRegion  Region = iota
	WetRegion    Region = iota
	NarrowRegion Region = iota
)

type Item uint8

const (
	ClimbingGearItem Item = iota
	TorchItem        Item = iota
	NeitherItem      Item = iota
)

var Dirs = []math.Vector2[int]{{X: -1}, {X: 1}, {Y: -1}, {Y: 1}}

var AllowedItems = map[Region]map[Item]bool{
	RockyRegion:  {ClimbingGearItem: true, TorchItem: true},
	WetRegion:    {ClimbingGearItem: true, NeitherItem: true},
	NarrowRegion: {TorchItem: true, NeitherItem: true},
}

type Cave struct {
	depth             int
	target            math.Vector2[int]
	geologicalIndeces map[math.Vector2[int]]int
	geologicalLevels  map[math.Vector2[int]]int
	regions           map[math.Vector2[int]]Region
}

func NewCave(depth int, target math.Vector2[int]) Cave {
	return Cave{
		depth:             depth,
		target:            target,
		geologicalIndeces: map[math.Vector2[int]]int{},
		geologicalLevels:  map[math.Vector2[int]]int{},
		regions:           map[math.Vector2[int]]Region{},
	}
}

func (c Cave) erosionLevel(pos math.Vector2[int]) int {
	if el, ok := c.geologicalLevels[pos]; ok {
		return el
	}

	gi := c.geologicalIndex(pos)
	erosionLevel := (gi + c.depth) % 20183
	c.geologicalLevels[pos] = erosionLevel

	return erosionLevel
}

func (c Cave) geologicalIndex(pos math.Vector2[int]) (gi int) {
	if gi, ok := c.geologicalIndeces[pos]; ok {
		return gi
	}

	defer func() { c.geologicalIndeces[pos] = gi }()

	if (pos.X == 0 && pos.Y == 0) || pos == c.target {
		return 0
	}

	if pos.Y == 0 {
		return pos.X * 16807
	}

	if pos.X == 0 {
		return pos.Y * 48271
	}

	return c.erosionLevel(pos.Add(math.NewVector2(-1, 0))) * c.erosionLevel(pos.Add(math.NewVector2(0, -1)))
}

func (c Cave) Region(pos math.Vector2[int]) Region {
	if r, ok := c.regions[pos]; ok {
		return r
	}

	region := Region(c.erosionLevel(pos) % 3)
	c.regions[pos] = region

	return region
}
